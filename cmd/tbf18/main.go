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
		fmt.Printf("CSVファイルを開けませんでした: %v\n", err)
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
		fmt.Printf("xlsxファイルを開けませんでした: %v\n", err)
		exitCode = 1
		return
	}
	defer xlsxFile.Close()

	// XLSX仕訳IOインスタンスを作成
	仕訳一覧xlsx := io.New仕訳XlsxIo(xlsxFile)

	// Service仕訳インスタンスを作成
	仕訳サービス := domain.NewService仕訳(csvReader, 仕訳一覧xlsx)

	// Service仕訳を実行
	仕訳一覧, err := 仕訳サービス.Query仕訳一覧()
	if err != nil && err != domain.Error未定義仕訳 {
		fmt.Printf("仕訳処理エラー: %v\n", err)
		exitCode = 1
		return
	} else if err == domain.Error未定義仕訳 {
		仕訳一覧xlsx.Save(仕訳一覧)
		exitCode = 2
		return
	}
	err = 仕訳一覧xlsx.Save(仕訳一覧)
	if err != nil {
		fmt.Printf("仕訳データの保存エラー: %v\n", err)
		exitCode = 2
		return
	}

	集計仕訳一覧, err := 仕訳サービス.Query集計仕訳(仕訳一覧)
	if err != nil {
		fmt.Printf("集計仕訳処理エラー: %v\n", err)
		exitCode = 2
		return
	}

	集計仕訳一覧xlsx := io.New集計仕訳XlsxWriter(xlsxFile)
	err = 集計仕訳一覧xlsx.Save(集計仕訳一覧.Get())
	if err != nil {
		fmt.Printf("集計仕訳データの保存エラー: %v\n", err)
		exitCode = 2
		return
	}

	// XLSX仕訳IOインスタンスを作成
	按分ルール一覧xlsx := io.New按分ルールXlsxIo(xlsxFile)
	// XLSX按分結果明細ライターインスタンスを作成
	按分結果明細xlsx := io.New按分結果明細XlsxWriter(xlsxFile)

	// Service配賦インスタンスを作成
	配賦サービス := domain.NewService配賦(按分ルール一覧xlsx, 按分結果明細xlsx)
	按分ルール一覧, err := 配賦サービス.Query按分ルール一覧()
	if err != nil {
		fmt.Printf("按分ルール処理エラー: %v\n", err)
		exitCode = 2
		return
	}
	按分結果明細一覧, err := 配賦サービス.Execute配賦(集計仕訳一覧.Get(), 按分ルール一覧)
	if err != nil {
		fmt.Printf("按分ルール処理エラー: %v\n", err)
		exitCode = 2
		return
	}
	for _, e := range 按分結果明細一覧 {
		fmt.Printf("%+v\n", e)
	}
}
