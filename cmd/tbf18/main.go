package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"

	"teckbookfest18-sample/domain"
	"teckbookfest18-sample/io"
)

func main() {
	exitCode := 0
	defer func() {
		if exitCode != 0 {
			fmt.Println("異常終了しました。")
		}
		os.Exit(exitCode)
	}()

	// CSVファイルを開く
	csvFile, err := os.Open("./sample_data.csv")
	if err != nil {
		fmt.Printf("CSVファイルを開けませんでした: %v", err)
		exitCode = 1
		return
	}
	defer csvFile.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true // 引用符の処理を緩和
	reader.Comma = ','       // 区切り文字を指定

	// CSV仕訳リーダーインスタンスを作成
	csvReader := io.New仕訳CsvReader(reader)

	// xlsxファイルを開く
	xlsxFile, err := excelize.OpenFile("./按分サンプル.xlsx") // xlsxファイルのパスを適切に設定してください
	if err != nil {
		fmt.Printf("xlsxファイルを開けませんでした: %v", err)
		exitCode = 1
		return
	}
	defer xlsxFile.Close()

	// XLSX仕訳IOインスタンスを作成
	xlsx仕訳一覧 := io.New仕訳XlsxIo(xlsxFile)

	// Service仕訳インスタンスを作成
	service := domain.NewService仕訳(csvReader, xlsx仕訳一覧)

	// Service仕訳を実行
	仕訳一覧, err := service.Query仕訳一覧()
	if err != nil && err != domain.Error未定義仕訳 {
		fmt.Printf("仕訳処理エラー: %v", err)
		exitCode = 1
		return
	} else if err == domain.Error未定義仕訳 {
		xlsx仕訳一覧.Save(仕訳一覧)
		exitCode = 2
		return
	}
	xlsx仕訳一覧.Save(仕訳一覧)

	集計仕訳一覧, err := service.Query集計仕訳(仕訳一覧)
	if err != nil {
		exitCode = 2
		return
	}

	xlsx集計仕訳一覧 := io.New集計仕訳XlsxWriter(xlsxFile)
	err = xlsx集計仕訳一覧.Save(集計仕訳一覧.Get())
	if err != nil {
		exitCode = 2
		return
	}

}
