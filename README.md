# 会計按分・仕訳処理ツール

このリポジトリは、会計業務における按分計算や仕訳処理を自動化するためのツールセットです。

## 機能概要

- 仕訳データのCSV/Excelファイル読み込み
- 集計仕訳の作成と出力
- 勤務表データからの工数集計（別ツールmhgによる処理）
- 按分ルールに基づく按分計算
- 按分結果のExcel出力

## プロジェクト構成

```
cmd/
  mhg/      - 勤務表処理用コマンド
  tbf18/    - 按分計算用コマンド
domain/     - ドメインモデルとビジネスロジック
io/         - ファイル入出力の実装
```

## 使い方

### 按分計算 (tbf18)

```bash
# デフォルトファイルを使用する場合
cd cmd/tbf18
go run main.go

# CSVファイルとExcelファイルを指定する場合
cd cmd/tbf18
go run main.go data.csv workbook.xlsx

# ヘルプを表示
cd cmd/tbf18
go run main.go -help
```

按分計算を実行します。
引数を指定しない場合は、サンプルデータとして`sample_data.csv`と`按分サンプル.xlsx`を使用します。

### 勤務表処理 (mhg)

```bash
cd cmd/mhg
go run main.go
```

勤務表から工数集計を行います。

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
# Windows環境での実行例
tbf18.exe                        # デフォルトファイルを使用
tbf18.exe data.csv workbook.xlsx # 独自のファイルパスを指定
tbf18.exe -help                  # ヘルプを表示
```

## 開発環境

- Go 言語<img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic" alt="Go">
- 依存ライブラリは`go.mod`ファイルに記載

## CREDITSファイル生成方法

- 初回のみ
  - `go install github.com/Songmu/gocredits/cmd/gocredits@latest`
- go.modを`go mod tidy`で最新化したら
  - `gocredits . > CREDITS`

## 貢献について

プルリクエストやイシューの報告は歓迎します。大きな変更を加える場合は、まずイシューで議論を開始してください。

## 注意事項

このツールは会計業務の効率化を目的としていますが、最終的な会計処理の責任はユーザーにあります。
出力結果は必ず確認してください。

## ライセンス

Copyright 2025 Moiterika LLC.  
Licensed under the MIT License.
詳細は[LICENSE](LICENSE)ファイルをご覧ください。