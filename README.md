# 配賦ツール、工数集計ツール

このリポジトリは、配賦や工数集計を自動化するためのツールセットです。
技術書典18で頒布の『原価管理・按分・Go』のサンプルです。

## 機能概要

- 仕訳データのCSV/Excelファイル読み込み
- 集計仕訳の作成と出力
- 勤務表データからの工数集計（別ツールmhgによる処理）
- 按分ルールに基づく按分計算
- 按分結果のExcel出力

## プロジェクト構成

```
cmd/
  mhg/      - 工数集計ツール（勤務表処理用コマンド）
  tbf18/    - 配賦ツール（按分計算用コマンド）
domain/     - ドメインモデルとビジネスロジック
io/         - ファイル入出力の実装
```

## 使い方

### 配賦ツール (tbf18)

```bash
# デフォルトファイルを使用する場合
cd cmd/tbf18
go run main.go

# CSVファイルとExcelファイルを指定する例
cd cmd/tbf18
go run main.go data.csv workbook.xlsx

# ヘルプを表示
cd cmd/tbf18
go run main.go -help
```

按分計算を実行します。
引数を指定しない場合は、サンプルデータとして`sample_data.csv`と`按分サンプル.xlsx`を使用します。

### 工数集計ツール (mhg)

```bash
# デフォルトファイルを使用する場合
cd cmd/mhg
go run main.go

# 勤務表ファイルと出力ファイルを指定する場合
cd cmd/mhg
go run main.go 勤務表.xlsx 出力.xlsx

# ヘルプを表示
cd cmd/mhg
go run main.go -help
```

勤務表から工数集計を行います。
引数を指定しない場合は、デフォルトで`勤務表_2024.xlsx`と`../tbf18/按分サンプル.xlsx`を使用します。

## ビルド方法

### Linux上でのWindows向けクロスコンパイル

```bash
# 按分計算ツール (Windows向け)
cd cmd/tbf18
GOOS=windows GOARCH=amd64 go build -o tbf18.exe

# 勤務表処理ツール (Windows向け)
cd cmd/mhg
GOOS=windows GOARCH=amd64 go build -o ../tbf18/mhg.exe
```

生成された `.exe` ファイルはWindows環境で実行できます。

```bash
# Windows環境での実行例（tbf18）
tbf18.exe                        # デフォルトファイルを使用
tbf18.exe data.csv workbook.xlsx # 独自のファイルパスを指定
tbf18.exe -help                  # ヘルプを表示

# Windows環境での実行例（mhg）
mhg.exe                          # デフォルトファイルを使用
mhg.exe 勤務表.xlsx 出力.xlsx     # 独自のファイルパスを指定
mhg.exe -help                    # ヘルプを表示
```

## 開発環境

- Go 言語<img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic" alt="Go">
- 依存ライブラリは`go.mod`ファイルに記載

## （筆者向け）CREDITSファイル生成方法

- 初回のみ
  - `go install github.com/Songmu/gocredits/cmd/gocredits@latest`
- go.modを`go mod tidy`で最新化したら
  - `gocredits . > CREDITS`

## 注意事項

このツールは会計業務の効率化を目的としていますが、最終的な会計処理の責任はユーザーにあります。
出力結果は必ず確認してください。  

developブランチは書籍執筆時からバージョンが進んでいる可能性があります。  

## ライセンス

Copyright 2025 Moiterika LLC.  
Licensed under the MIT License.
詳細は[LICENSE](LICENSE)ファイルをご覧ください。