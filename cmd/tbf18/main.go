package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"teckbookfest18-sample/io"
)

func main() {
	// SQLite3データベースを開く（存在しない場合は作成）
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalf("データベースを開けませんでした: %v", err)
	}
	defer db.Close()

	// SQLite仕訳Repoインスタンスを作成
	repo := io.NewSQLite仕訳Repo(db)

	// 仕訳テーブルを作成
	err = repo.CreateTableIfNotExists()
	if err != nil {
		log.Fatalf("仕訳テーブル群作成エラー: %v", err)
	}

	// CSVファイルを開く - 実際のCSVファイルパスに置き換えてください
	// 注意: この例ではサンプルのCSVファイルがあることを前提としています
	csvFile, err := os.Open("./sample_data.csv")
	if err != nil {
		log.Printf("CSVファイルを開けませんでした: %v", err)
		fmt.Println("SQLite3データベースに接続しました！")
		return
	}
	defer csvFile.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true // 引用符の処理を緩和
	reader.Comma = ','       // 区切り文字を指定

	// Query仕訳インスタンスを作成
	query := io.NewQuery仕訳Csv(reader)

	// CSVデータを読み取る
	仕訳一覧, err := query.Read()
	if err != nil {
		log.Printf("CSVデータの読み取りエラー: %v", err)
		fmt.Println("SQLite3データベースに接続しました！")
		return
	}

	// 読み取ったデータをデータベースに保存
	err = repo.Save(仕訳一覧)
	if err != nil {
		log.Fatalf("データベース保存エラー: %v", err)
	}

	fmt.Printf("SQLite3データベースに接続し、%d件の仕訳データを保存しました！\n", len(仕訳一覧))

	// 保存したデータの取得テスト
	savedData, err := repo.FindAll()
	if err != nil {
		log.Printf("データ取得エラー: %v", err)
	} else {
		fmt.Printf("データベースから%d件の仕訳データを取得しました。\n", len(savedData))
	}
}
