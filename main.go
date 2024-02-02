// main.go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// コマンドラインオプションの定義
	var seed string
	flag.StringVar(&seed, "seed", "", "乱数のシード値を指定")
	var spreadsheetFile string
	flag.StringVar(&spreadsheetFile, "xlsx", "spreadsheet.xlsx", "読み込むスプレッドシートファイルのパス")
	// 他のフラグと共にパースする
	flag.Parse()

	// スプレッドシートからデータを取得
	data, err := FetchSheetsData(spreadsheetFile)
	if err != nil {
		fmt.Printf("スプレッドシートからデータを取得できませんでした: %v\n", err)
		os.Exit(1)
	}

	// クイズを生成
	quiz, err := GenerateQuiz(data, 10, seed)
	if err != nil {
		fmt.Printf("クイズを生成できませんでした: %v\n", err)
		os.Exit(1)
	}
	// クイズを実行
	// テスト環境ではクイズを実行しない
	if os.Getenv("APP_ENV") != "test" {
		// クイズを実行
		RunQuiz(quiz)

		return
	}
}
