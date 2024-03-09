package sheets

import (
	"fmt"
	"io/fs"

	"github.com/xuri/excelize/v2"
)

// SheetData はスプレッドシートのデータを保持する構造体です。
type SheetData struct {
	Number      string
	Word        string
	Translation string
}

// FetchSheetsData はファイルパスからExcelスプレッドシートのデータを取得します。
func FetchSheetsData(filePath string) ([]SheetData, error) {
	excelFile, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer excelFile.Close()

	return readSheetData(excelFile)
}

// FetchSheetsDataFromFS は embed.FS からExcelスプレッドシートのデータを取得します。
func FetchSheetsDataFromFS(fsys fs.FS, filePath string) ([]SheetData, error) {
	f, err := fsys.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file from embedded FS: %w", err)
	}
	defer f.Close()

	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read Excel file from embedded FS: %w", err)
	}
	defer excelFile.Close()

	return readSheetData(excelFile)
}

// readSheetData はExcelファイルからシートデータを読み込みます。
func readSheetData(excelFile *excelize.File) ([]SheetData, error) {
	var data []SheetData
	sheets := excelFile.GetSheetMap()
	for _, sheetName := range sheets {
		rows, err := excelFile.GetRows(sheetName)
		if err != nil {
			return nil, fmt.Errorf("failed to get rows from sheet '%s': %w", sheetName, err)
		}

		for i, row := range rows {
			if i == 0 || row[0] == "#" {
				continue // ヘッダー行またはコメント行をスキップ
			}

			if len(row) < 3 {
				return nil, fmt.Errorf("invalid row format at sheet '%s', row %d", sheetName, i+1)
			}

			data = append(data, SheetData{
				Number:      row[0],
				Word:        row[1],
				Translation: row[2],
			})
		}
	}

	return data, nil
}
