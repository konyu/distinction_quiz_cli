package quiz

import (
	"distinction_quiz_cli/sheets"
	"distinction_quiz_cli/utils"
	"fmt"
)

// QuizItem represents a single quiz question.
type QuizItem struct {
	Word        string
	Translation string
	Choices     []string
}

// GenerateQuiz creates a quiz from the spreadsheet data.
func GenerateQuiz(data []sheets.SheetData, numQuestions int, seed string) ([]QuizItem, error) {
	r := utils.GenerateRand(seed)
	utils.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] }, r)

	if len(data) < numQuestions {
		return nil, fmt.Errorf("not enough data to generate quiz: available %d, required %d", len(data), numQuestions)
	}

	quiz := make([]QuizItem, numQuestions)
	for i := 0; i < numQuestions; i++ {
		correct := data[i]
		choices := []string{correct.Translation}

		// Add incorrect choices
		for len(choices) < 4 {
			j := r.Intn(len(data))
			if j != i && !utils.Contains(choices, data[j].Translation) {
				choices = append(choices, data[j].Translation)
			}
		}

		utils.Shuffle(len(choices), func(i, j int) { choices[i], choices[j] = choices[j], choices[i] }, r)
		quiz[i] = QuizItem{
			Word:        correct.Word,
			Translation: correct.Translation,
			Choices:     choices,
		}
	}

	return quiz, nil
}

// RunQuiz executes the quiz and processes user input.
func RunQuiz(quizItems []QuizItem) {
	var correctCount int

	utils.ClearScreen()

	for i, q := range quizItems {
		displayQuestion(i, q)
		userChoice := getUserChoice()
		utils.ClearScreen()
		correct := checkAnswer(q, userChoice)
		displayAnswerResult(correct, q)

		if correct {
			correctCount++
		}
		displaySeparator()
	}

	displayFinalResult(correctCount, len(quizItems))
}

// displayQuestion prints the quiz question.
func displayQuestion(index int, q QuizItem) {
	fmt.Printf("Question %d: %s\n\n", index+1, q.Word)
	for idx, choice := range q.Choices {
		fmt.Printf("%d: %s\n", idx+1, choice)
	}
}

// getUserChoice gets the user's answer.
func getUserChoice() int {
	var userChoice int
	fmt.Printf("\nPlease enter your answer (1-4): ")
	_, err := fmt.Scan(&userChoice)
	if err != nil || userChoice < 1 || userChoice > 4 {
		fmt.Println("Enter a number between 1 and 4.")
		return getUserChoice()
	}
	return userChoice
}

// checkAnswer checks if the user's answer is correct.
func checkAnswer(q QuizItem, userChoice int) bool {
	return q.Choices[userChoice-1] == q.Translation
}

// displayAnswerResult prints the result of the answer.
func displayAnswerResult(correct bool, q QuizItem) {
	explanation := fmt.Sprintf("\"%s\" means \"%s\".", q.Word, q.Translation)
	if correct {
		fmt.Printf("\033[32mCorrect! \033[0m%s\n", explanation)
	} else {
		fmt.Printf("\033[31mIncorrect.\033[0m %s\n", explanation)
	}
}

// displaySeparator prints a separator line.
func displaySeparator() {
	fmt.Println("\n--------")
}

// displayFinalResult prints the final quiz results.
func displayFinalResult(correctCount, totalQuestions int) {
	fmt.Printf("In this %d-question quiz, you got %d right.\n", totalQuestions, correctCount)
}
