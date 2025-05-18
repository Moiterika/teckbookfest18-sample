package io

import (
	"teckbookfest18-sample/domain"

	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
)

// 勤務表XlsxReader は Qry勤務表 の実装です
// 勤務表データのxlsxファイル読み取りを担当します
type 勤務表XlsxReader struct {
	ef *excelize.File
}

// New勤務表XlsxReader は勤務表XlsxReaderを生成します
func New勤務表XlsxReader(ef *excelize.File) *勤務表XlsxReader {
	return &勤務表XlsxReader{ef: ef}
}

const sheet勤務表 = "勤務表"

// 勤務表のExcelシートの列インデックス定義
const (
	FldIdx作業内容       = 5
	FldIdx作業時間_分     = 7
	FldIdx労務費按分用の計上月 = 8
	FldIdx経費按分用の計上月  = 9
)

// Read勤務表 は勤務表データを読み取ります
func (x *勤務表XlsxReader) Read勤務表() ([]*domain.Ent勤務表, error) {
	rows, err := x.ef.GetRows(sheet勤務表)
	if err != nil {
		return make([]*domain.Ent勤務表, 0), err
	}

	ret := make([]*domain.Ent勤務表, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			// ヘッダーを読み飛ばし
			continue
		}

		// 行データのチェック
		if len(row) <= FldIdx経費按分用の計上月 {
			// 必要なデータが不足している行はスキップ
			continue
		}

		作業時間_分, err := getDecimalCell(row, FldIdx作業時間_分)
		if err != nil {
			// 作業時間が取得できない場合はゼロ値を使用
			作業時間_分 = decimal.NewFromInt(0)
		}

		var e domain.Ent勤務表
		e.Fld作業内容 = getStringCell(row, FldIdx作業内容)
		e.Fld作業時間_分 = 作業時間_分
		e.Fld労務費按分用の計上月 = getStringCell(row, FldIdx労務費按分用の計上月)
		e.Fld経費按分用の計上月 = getStringCell(row, FldIdx経費按分用の計上月)

		ret = append(ret, &e)
	}

	return ret, nil
}
