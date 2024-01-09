// quiz.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// QuizItem はクイズの一問を表す構造体です。
type QuizItem struct {
	Word        string
	Translation string
	Choices     []string
}

// GenerateQuiz はスプレッドシートのデータからクイズを生成します。
func GenerateQuiz(data []SheetData, numQuestions int) ([]QuizItem, error) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
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
		fmt.Printf("問題 %d: %s\n\n", i+1, q.Word)
		for idx, choice := range q.Choices {
			fmt.Printf("%d: %s\n", idx+1, choice)
		}

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
		// # CLIの画面をclearする
		fmt.Print("\033[H\033[2J")
		explanation := fmt.Sprintf("\"%s\" の意味は \"%s\" です。", q.Word, q.Translation)

		if q.Choices[userChoice-1] == q.Translation {
			// コンソールの文字の色を緑にする
			fmt.Printf("\033[32m⭕ 正解です! \033[0m%s\n", explanation)
			correctCount++
		} else {
			fmt.Printf("\033[31m❌ 不正解です。\033[0m%s\n", explanation)
		}
		fmt.Println("\n--------")
	}

	fmt.Printf("正解数 %d/%d\n", correctCount, len(quiz))
}
