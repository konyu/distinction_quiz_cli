package main

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"distinction_quiz_cli/quiz"
	"distinction_quiz_cli/sheets"
)

// スプレッドシートを埋め込む

//go:embed spreadsheet.xlsx
var embeddedFiles embed.FS

func main() {
	// コマンドラインオプションの定義
	var seed string
	flag.StringVar(&seed, "seed", "", "乱数のシード値を指定")
	var spreadsheetFile string
	flag.StringVar(&spreadsheetFile, "xlsx", "", "読み込むスプレッドシートファイルのパス")
	// 他のフラグと共にパースする
	flag.Parse()

	// スプレッドシートからデータを取得
	data := getSpreadSheetData(spreadsheetFile)
	// クイズを生成
	quizItems, err := quiz.GenerateQuiz(data, 10, seed)
	if err != nil {
		fmt.Printf("クイズを生成できませんでした: %v\n", err)
		os.Exit(1)
	}
	// クイズを実行
	// テスト環境ではクイズを実行しない
	if os.Getenv("APP_ENV") != "test" {
		// クイズを実行
		quiz.RunQuiz(quizItems)

		return
	}
}

func getSpreadSheetData(spreadsheetFile string) []sheets.SheetData {
	var data []sheets.SheetData
	var err error
	if spreadsheetFile == "" {

		data, err = sheets.FetchSheetsDataFromFS(embeddedFiles, "spreadsheet.xlsx")
		if err != nil {
			fmt.Printf("スプレッドシートからデータを取得できませんでした: %v\n", err)
			os.Exit(1)
		}
	} else {
		data, err = sheets.FetchSheetsData(spreadsheetFile)
		if err != nil {
			fmt.Printf("スプレッドシートからデータを取得できませんでした: %v\n", err)
		}
	}
	return data
}
