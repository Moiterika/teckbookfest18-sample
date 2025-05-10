package io

import (
	"fmt"
	"teckbookfest18-sample/domain"

	"github.com/xuri/excelize/v2"
)

// 集約仕訳XlsxWriter は I集約仕訳XlsxWriter の実装です
// 集計仕訳データをxlsxファイルに書き出します
type 集約仕訳XlsxWriter struct {
	ef *excelize.File
}

// New集約仕訳XlsxWriter は Excelize ファイルを受け取り、新しい Writer を返します
func New集約仕訳XlsxWriter(ef *excelize.File) *集約仕訳XlsxWriter {
	return &集約仕訳XlsxWriter{ef: ef}
}

const sheet集約仕訳一覧 = "集約仕訳一覧"

// Save は集計仕訳一覧をxlsxに書き出します
func (w *集約仕訳XlsxWriter) Save(data []*domain.Ent集計仕訳) error {
	// 既にシートが存在する場合は削除
	if _, err := w.ef.GetSheetIndex(sheet集約仕訳一覧); err == nil {
		if err := w.ef.DeleteSheet(sheet集約仕訳一覧); err != nil {
			return err
		}
	}

	// シートを作成
	idx, err := w.ef.NewSheet(sheet集約仕訳一覧)
	if err != nil {
		return err
	}
	// ヘッダー行を書き込み
	headers := []interface{}{"計上年月", "コストプール", "按分ルール1", "按分ルール2", "借方税区分", "借方税率", "合計金額"}
	w.ef.SetSheetRow(sheet集約仕訳一覧, "A1", &headers)
	// データ行を書き込み
	for i, e := range data {
		row := []interface{}{e.Fld計上年月, e.Fldコストプール, e.Fld按分ルール1, e.Fld按分ルール2, e.Fld借方税区分, e.Fld借方税率.String(), e.Fld合計金額.String()}
		cell := fmt.Sprintf("A%d", i+2)
		w.ef.SetSheetRow(sheet集約仕訳一覧, cell, &row)
	}
	// 新規シートをアクティブに設定
	w.ef.SetActiveSheet(idx)
	return nil
}
