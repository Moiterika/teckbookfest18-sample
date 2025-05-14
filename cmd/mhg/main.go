package main

import (
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
		fmt.Println("使用方法: ./mhg [勤務表ファイルパス] [出力Excelファイルパス]")
		fmt.Println("デフォルト値:")
		fmt.Println("  勤務表ファイルパス: ./勤務表_2024.xlsx")
		fmt.Println("  出力Excelファイルパス: ../tbf18/按分サンプル.xlsx")
		fmt.Println("\nオプション:")
		flag.PrintDefaults()
		return
	}

	// 位置引数の処理
	args := flag.Args()
	勤務表FilePath := "./勤務表_2024.xlsx"         // デフォルト値
	outputFilePath := "../tbf18/按分サンプル.xlsx" // デフォルト値

	// 引数の数に応じて処理
	if len(args) >= 1 {
		勤務表FilePath = args[0]
	}
	if len(args) >= 2 {
		outputFilePath = args[1]
	}

	exitCode := 0
	defer func() {
		if exitCode != 0 {
			fmt.Println("異常終了しました。")
		}
		os.Exit(exitCode)
	}()
	// xlsxファイルを開く
	勤務表xlsx, err := excelize.OpenFile(勤務表FilePath)
	if err != nil {
		fmt.Printf("勤務表ファイル '%s' を開けませんでした: %v\n", 勤務表FilePath, err)
		exitCode = 1
		return
	}
	defer 勤務表xlsx.Close()

	outputXlsx, err := excelize.OpenFile(outputFilePath)
	if err != nil {
		fmt.Printf("出力Excelファイル '%s' を開けませんでした: %v\n", outputFilePath, err)
		exitCode = 1
		return
	}
	defer outputXlsx.Close()

	// リーダーとライターのインスタンスを作成
	勤務表Reader := io.New勤務表XlsxReader(勤務表xlsx)
	按分ルールIo := io.New按分ルールXlsxIo(outputXlsx)

	// Service工数集計インスタンスを作成
	工数集計サービス := domain.NewService工数集計(勤務表Reader, 按分ルールIo)

	// 工数集計処理を実行
	err = 工数集計サービス.Execute工数集計()
	if err != nil {
		fmt.Printf("工数集計処理に失敗: %v\n", err)
		exitCode = 1
		return
	}

}
