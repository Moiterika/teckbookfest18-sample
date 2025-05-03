package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// SQLite3データベースを開く（存在しない場合は作成）
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalf("データベースを開けませんでした: %v", err)
	}
	defer db.Close()

	// テーブル作成
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER
		)
	`)
	if err != nil {
		log.Fatalf("テーブル作成エラー: %v", err)
	}

	fmt.Println("SQLite3データベースに接続しました！")
}
