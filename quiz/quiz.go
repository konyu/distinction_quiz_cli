package quiz

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"time"

	"distinction_quiz_cli/sheets"
)

// QuizItem はクイズの一問を表す構造体です。
type QuizItem struct {
	Word        string
	Translation string
	Choices     []string
}

// GenerateQuiz はスプレッドシートのデータからクイズを生成します。
func GenerateQuiz(data []sheets.SheetData, numQuestions int, seed string) ([]QuizItem, error) {
	r := generateRand(seed)
	r.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	if len(data) < numQuestions {
		return nil, fmt.Errorf("not enough data to generate quiz")
	}

	var quiz []QuizItem
	for i := 0; i < numQuestions; i++ {
		correct := data[i]
		choices := make([]string, 1)
		choices[0] = correct.Translation

		// 正解以外の選択肢を追加
		for len(choices) < 4 {
			j := rand.Intn(len(data))
			if j != i && !contains(choices, data[j].Translation) {
				choices = append(choices, data[j].Translation)
			}
		}

		r.Shuffle(len(choices), func(i, j int) { choices[i], choices[j] = choices[j], choices[i] })
		quiz = append(quiz, QuizItem{
			Word:        correct.Word,
			Translation: correct.Translation,
			Choices:     choices,
		})
	}

	return quiz, nil
}

// contains はスライスに特定の文字列が含まれているかをチェックします。
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// RunQuiz はクイズを実行し、ユーザーの入力を処理します。
func RunQuiz(quiz []QuizItem) {
	var correctCount int
	for i, q := range quiz {
		displayQuestion(i, q)
		userChoice := getUserChoice()

		clearScreen()
		correct := checkAnswer(q, userChoice)
		displayAnswerResult(correct, q)

		if correct {
			correctCount++
		}
		displaySeparator()
	}

	displayFinalResult(correctCount, len(quiz))
}

// displayQuestion はクイズの問題を表示します。
func displayQuestion(index int, q QuizItem) {
	fmt.Printf("問題 %d: %s\n\n", index+1, q.Word)
	for idx, choice := range q.Choices {
		fmt.Printf("%d: %s\n", idx+1, choice)
	}
}

// getUserChoice はユーザーの回答を取得します。
func getUserChoice() int {
	var userChoice int
	for {
		fmt.Printf("\n回答してください (1-4): ")
		_, err := fmt.Scan(&userChoice)
		if err != nil || userChoice < 1 || userChoice > 4 {
			fmt.Println("1から4までの数値を入力してください。")
			continue
		}
		break
	}
	return userChoice
}

// clearScreen はCLIの画面をクリアします。
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// checkAnswer はユーザーの回答が正しいかどうかをチェックします。
func checkAnswer(q QuizItem, userChoice int) bool {
	return q.Choices[userChoice-1] == q.Translation
}

// displayAnswerResult は回答の結果を表示します。
func displayAnswerResult(correct bool, q QuizItem) {
	explanation := fmt.Sprintf("\"%s\" の意味は \"%s\" です。", q.Word, q.Translation)
	if correct {
		fmt.Printf("\033[32m⭕ 正解です! \033[0m%s\n", explanation)
	} else {
		fmt.Printf("\033[31m❌ 不正解です。\033[0m%s\n", explanation)
	}
}

// displaySeparator は区切り線を表示します。
func displaySeparator() {
	fmt.Println("\n--------")
}

// displayFinalResult は最終結果を表示します。
func displayFinalResult(correctCount, totalQuestions int) {
	fmt.Printf("正解数 %d/%d\n", correctCount, totalQuestions)
}

func generateRand(seed string) *rand.Rand {
	var r *rand.Rand

	if seed != "" {
		h := fnv.New64a()
		_, err := h.Write([]byte(seed))
		if err != nil {
			fmt.Printf("シード値のハッシュ生成中にエラーが発生しました: %v\n", err)
			os.Exit(1)
		}
		seedInt := int64(h.Sum64())
		r = rand.New(rand.NewSource(seedInt))
	} else {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return r
}
