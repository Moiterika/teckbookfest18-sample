package main

import (
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
	// xlsxファイルを開く
	勤務表xlsx, err := excelize.OpenFile("./勤務表_2024.xlsx") // パスは必要に応じて調整
	if err != nil {
		fmt.Printf("xlsxファイルを開けませんでした: %v\n", err)
		exitCode = 1
		return
	}
	defer 勤務表xlsx.Close()

	outputXlsx, err := excelize.OpenFile("../tbf18/按分サンプル.xlsx") // パスは必要に応じて調整
	if err != nil {
		fmt.Printf("xlsxファイルを開けませんでした: %v\n", err)
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
