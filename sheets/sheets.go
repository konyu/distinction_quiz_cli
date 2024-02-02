package sheets

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// SheetData はスプレッドシートのデータを保持する構造体です。
type SheetData struct {
	Number      string
	Word        string
	Translation string
}

var demoFlg string = ""

// FetchSheetsData は指定されたファイルのExcelスプレッドシートからデータを取得します。
func FetchSheetsData(filePath string) ([]SheetData, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %v", err)
	}
	defer f.Close()

	var data []SheetData
	for _, sheetName := range f.GetSheetMap() {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, fmt.Errorf("failed to get rows from sheet '%s': %v", sheetName, err)
		}

		for i, row := range rows {
			// ヘッダー行をスキップ
			if i == 0 || row[0] == "#" {
				continue
			}

			if demoFlg == "true" && i > 21 {
				break
			}
			// 行のデータが足りない場合はスキップ
			if len(row) < 3 {
				continue
			}
			// SheetDataにデータを追加
			data = append(data, SheetData{
				Number:      row[0],
				Word:        row[1],
				Translation: row[2],
			})
		}
	}
	if demoFlg == "true" {
		fmt.Println("DEMO MODE!!!!")
	}

	return data, nil
}
