package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"

	"teckbookfest18-sample/domain"
	"teckbookfest18-sample/io"
)

func main() {
	// ヘルプフラグの定義（オプション引数として残す）
	helpFlag := flag.Bool("help", false, "ヘルプを表示")
	flag.Parse()

	// ヘルプフラグが指定された場合、使用方法を表示して終了
	if *helpFlag {
		fmt.Println("使用方法: ./tbf18 [CSVファイルパス] [Excelファイルパス]")
		fmt.Println("デフォルト値:")
		fmt.Println("  CSVファイルパス: ./sample_data.csv")
		fmt.Println("  Excelファイルパス: ./按分サンプル.xlsx")
		fmt.Println("\nオプション:")
		flag.PrintDefaults()
		return
	}

	// 位置引数の処理
	args := flag.Args()
	csvFilePath := "./sample_data.csv" // デフォルト値
	xlsxFilePath := "./按分サンプル.xlsx"    // デフォルト値

	// 引数の数に応じて処理
	if len(args) >= 1 {
		csvFilePath = args[0]
	}
	if len(args) >= 2 {
		xlsxFilePath = args[1]
	}

	exitCode := 0
	defer func() {
		if exitCode != 0 {
			fmt.Println("異常終了しました。")
		}
		os.Exit(exitCode)
	}()

	// CSVファイルを開く
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Printf("CSVファイル '%s' を開けませんでした: %v\n", csvFilePath, err)
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
	xlsxFile, err := excelize.OpenFile(xlsxFilePath)
	if err != nil {
		fmt.Printf("xlsxファイル '%s' を開けませんでした: %v\n", xlsxFilePath, err)
		exitCode = 1
		return
	}
	defer xlsxFile.Close()

	// XLSX仕訳IOインスタンスを作成
	仕訳一覧xlsx := io.New仕訳XlsxIo(xlsxFile)

	// 勘定科目Readerインスタンスを作成
	勘定科目xlsx := io.New勘定科目XlsxReader(xlsxFile)

	// Service仕訳インスタンスを作成
	仕訳サービス := domain.NewService仕訳(csvReader, 仕訳一覧xlsx, 勘定科目xlsx)

	// Service仕訳を実行
	仕訳一覧, err := 仕訳サービス.Query仕訳一覧()
	if err != nil {
		fmt.Printf("仕訳処理エラー: %v\n", err)
		exitCode = 1
		return
	}
	err = 仕訳サービス.Save(仕訳一覧)
	if errors.Is(err, domain.Error未定義仕訳) {
		fmt.Printf("仕訳一覧シートのA～E列に記入して下さい。\n")
		exitCode = 2
		return
	}
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
	// XLSX按分結果ライターインスタンスを作成
	按分結果xlsx := io.New按分結果XlsxWriter(xlsxFile)

	// Service配賦インスタンスを作成
	配賦サービス := domain.NewService配賦(按分ルール一覧xlsx, 按分結果明細xlsx, 按分結果xlsx)
	按分ルール一覧, err := 配賦サービス.Query按分ルール一覧()
	if err != nil {
		fmt.Printf("按分ルール処理エラー: %v\n", err)
		exitCode = 2
		return
	}
	err = 配賦サービス.Execute配賦(集計仕訳一覧.Get(), 按分ルール一覧)
	if err != nil {
		fmt.Printf("按分ルール処理エラー: %v\n", err)
		exitCode = 2
		return
	}
}
