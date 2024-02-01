// main.go
package main

import (
	"flag"
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"time"
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

	var r *rand.Rand
	// シード値が指定されていれば乱数のシードを設定
	if seed != "" {
		// 文字列からハッシュ値を生成し、それをシードとして使用
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

	// クイズを生成
	quiz, err := GenerateQuiz(data, 10, r)
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
