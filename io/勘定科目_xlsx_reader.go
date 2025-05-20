package io

import (
	"teckbookfest18-sample/domain"

	"github.com/xuri/excelize/v2"
)

// 勘定科目XlsxReader は Qry勘定科目 の実装です
// 勘定科目データのxlsxファイル読み取りを担当します
type 勘定科目XlsxReader struct {
	ef *excelize.File
}

func New勘定科目XlsxReader(ef *excelize.File) *勘定科目XlsxReader {
	return &勘定科目XlsxReader{ef: ef}
}

const sheet勘定科目一覧 = "勘定科目一覧"

func (x *勘定科目XlsxReader) Read勘定科目一覧() ([]*domain.Ent勘定科目, error) {
	rows, err := x.ef.GetRows(sheet勘定科目一覧)
	if err != nil {
		return make([]*domain.Ent勘定科目, 0), err
	}
	ret := make([]*domain.Ent勘定科目, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			// ヘッダーを読み飛ばし
			continue
		}
		var e domain.Ent勘定科目
		e.Fld勘定科目 = getStringCell(row, 0)
		e.Fld基本ルール = domain.New基本ルール(getStringCell(row, 1))
		e.Fld原価要素 = getStringCell(row, 2)
		e.Fldコストプール = getStringCell(row, 3)
		ret = append(ret, &e)
	}
	return ret, nil
}
