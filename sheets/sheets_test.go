package sheets

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/xuri/excelize/v2"
)

type SheetTestSuite struct {
	suite.Suite
	TestData     []SheetData
	TestFilePath string
}

// SetupTestは各テストの前に実行されます。
func (suite *SheetTestSuite) SetupTest() {
	suite.TestData = []SheetData{
		{Number: "1", Word: "apple", Translation: "りんご"},
		{Number: "2", Word: "banana", Translation: "バナナ"},
	}
	suite.TestFilePath = "./test.xlsx"
	createTestFile(suite.T(), suite.TestFilePath, suite.TestData)
}

// TearDownTestは各テストの後に実行されます。
func (suite *SheetTestSuite) TearDownTest() {
	err := os.Remove(suite.TestFilePath)
	suite.NoError(err)
}

// TestFetchSheetsDataはFetchSheetsData関数をテストします。
func (suite *SheetTestSuite) TestFetchSheetsData() {
	// テスト実行
	got, err := FetchSheetsData(suite.TestFilePath)
	suite.NoError(err)

	// 結果の検証
	suite.Equal(suite.TestData, got)
}

// TestSuiteを実行するためのヘルパー関数です。
func TestSheetTestSuite(t *testing.T) {
	suite.Run(t, new(SheetTestSuite))
}

// createTestFileはテスト用のExcelファイルを作成するヘルパー関数です。
func createTestFile(t *testing.T, filePath string, data []SheetData) {
	t.Helper()
	f := excelize.NewFile()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		t.Fatalf("Error creating new sheet: %v", err)
	}
	for i, d := range data {
		cellIndex := fmt.Sprintf("A%d", i+2)
		f.SetCellValue("Sheet1", cellIndex, d.Number)
		cellIndex = fmt.Sprintf("B%d", i+2)
		f.SetCellValue("Sheet1", cellIndex, d.Word)
		cellIndex = fmt.Sprintf("C%d", i+2)
		f.SetCellValue("Sheet1", cellIndex, d.Translation)
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs(filePath); err != nil {
		t.Fatalf("Unable to create test file: %s", err)
	}
}
