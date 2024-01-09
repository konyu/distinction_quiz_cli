// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	spreadsheetFile := "spreadsheet.xlsx"

	// スプレッドシートからデータを取得
	data, err := FetchSheetsData(spreadsheetFile)
	if err != nil {
		fmt.Printf("スプレッドシートからデータを取得できませんでした: %v\n", err)
		os.Exit(1)
	}

	// クイズを生成
	quiz, err := GenerateQuiz(data, 10)
	if err != nil {
		fmt.Printf("クイズを生成できませんでした: %v\n", err)
		os.Exit(1)
	}

	// クイズを実行
	RunQuiz(quiz)
}
