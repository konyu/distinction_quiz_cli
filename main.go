package main

import (
	"distinction_quiz_cli/quiz"
	"distinction_quiz_cli/sheets"
	"embed"
	"flag"
	"log"
)

//go:embed spreadsheet.xlsx
var spreadsheetFS embed.FS

const (
	defaultSpreadsheet  = "spreadsheet.xlsx"
	defaultNumQuestions = 10 // デフォルトの問題数
)

func main() {
	// コマンドラインオプションの定義
	var seed string
	flag.StringVar(&seed, "seed", "", "乱数のシード値を指定")
	var spreadsheetFile string
	flag.StringVar(&spreadsheetFile, "xlsx", defaultSpreadsheet, "読み込むスプレッドシートファイルのパス")
	var numQuestions int
	flag.IntVar(&numQuestions, "num", defaultNumQuestions, "生成するクイズの問題数")
	flag.Parse()

	var data []sheets.SheetData
	var err error
	// コマンドラインオプションで指定されたファイルパスがデフォルトと異なる場合は、そのファイルを使用
	if spreadsheetFile != defaultSpreadsheet {
		data, err = sheets.FetchSheetsData(spreadsheetFile)
	} else {
		// デフォルトのファイルパスが指定されている場合は、埋め込んだファイルを使用
		data, err = sheets.FetchSheetsDataFromFS(spreadsheetFS, defaultSpreadsheet)
	}
	if err != nil {
		log.Fatalf("Error fetching sheet data: %v\n", err)
	}
	// クイズを生成
	quizItems, err := quiz.GenerateQuiz(data, numQuestions, seed)
	if err != nil {
		log.Fatalf("Error generating quiz: %v\n", err)
	}

	// クイズを実行
	quiz.RunQuiz(quizItems)
}
