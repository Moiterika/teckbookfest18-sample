package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"

	"teckbookfest18-sample/domain"
	"teckbookfest18-sample/io"
)

func main() {
	// CSVファイルを開く
	csvFile, err := os.Open("./sample_data.csv")
	if err != nil {
		log.Fatalf("CSVファイルを開けませんでした: %v", err)
	}
	defer csvFile.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true // 引用符の処理を緩和
	reader.Comma = ','       // 区切り文字を指定

	// CSV仕訳リーダーインスタンスを作成
	csvReader := io.New仕訳CsvReader(reader)

	// Excelファイルを開く
	xlsxFile, err := excelize.OpenFile("./template.xlsx") // テンプレートファイルのパスを適切に設定してください
	if err != nil {
		log.Fatalf("Excelファイルを開けませんでした: %v", err)
	}
	defer xlsxFile.Close()

	// XLSX仕訳IOインスタンスを作成
	xlsxIo := io.New仕訳XlsxIo(xlsxFile)

	// Service仕訳インスタンスを作成
	service := domain.NewService仕訳(csvReader, xlsxIo)

	// Service仕訳を実行
	err = service.Execute()
	if err != nil {
		log.Fatalf("仕訳処理エラー: %v", err)
	}

	fmt.Println("仕訳データの処理が完了しました！")
}
