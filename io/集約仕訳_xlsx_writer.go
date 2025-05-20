package io

import (
	"fmt"
	"teckbookfest18-sample/domain"

	"github.com/xuri/excelize/v2"
)

// 集計仕訳XlsxWriter は Cmd集計仕訳 の実装です
// 集計仕訳データをxlsxファイルに書き出します
type 集計仕訳XlsxWriter struct {
	ef *excelize.File
}

// New集計仕訳XlsxWriter は Excelize ファイルを受け取り、新しい Writer を返します
func New集計仕訳XlsxWriter(ef *excelize.File) *集計仕訳XlsxWriter {
	return &集計仕訳XlsxWriter{ef: ef}
}

const sheet集計仕訳一覧 = "集計仕訳一覧"

// Save は集計仕訳一覧をxlsxに書き出します
func (w *集計仕訳XlsxWriter) Save(集計仕訳一覧 []*domain.Ent集計仕訳) error {
	// シートの有無チェックと既存行取得
	var sheetIdx int
	var existingRowsCount int
	if idx, err := w.ef.GetSheetIndex(sheet集計仕訳一覧); err != nil {
		return err
	} else if idx != -1 {
		sheetIdx = idx
		existingRows, _ := w.ef.GetRows(sheet集計仕訳一覧)
		existingRowsCount = len(existingRows)
	} else {
		var err error
		sheetIdx, err = w.ef.NewSheet(sheet集計仕訳一覧)
		if err != nil {
			return err
		}
	}
	// ヘッダー行を書き込み
	headers := []interface{}{"計上年月", "原価要素", "コストプール", "按分ルール1", "按分ルール2", "借方税区分", "借方税率", "合計金額"}
	w.ef.SetSheetRow(sheet集計仕訳一覧, "A1", &headers)
	// データ行を書き込み
	for i, e := range 集計仕訳一覧 {
		row := []interface{}{e.Fld計上年月, e.Fld原価要素, e.Fldコストプール, e.Fld按分ルール1, e.Fld按分ルール2, e.Fld借方税区分, e.Fld借方税率.IntPart(), e.Fld合計金額.IntPart()}
		cell := fmt.Sprintf("A%d", i+2)
		w.ef.SetSheetRow(sheet集計仕訳一覧, cell, &row)
	}
	// 余分な行を削除（データ数が減った場合）
	if existingRowsCount > len(集計仕訳一覧)+1 {
		// 余分な行を1行ずつ削除
		targetRowNum := len(集計仕訳一覧) + 2
		count := existingRowsCount - (len(集計仕訳一覧) + 1)
		for i := 0; i < count; i++ {
			if err := w.ef.RemoveRow(sheet集計仕訳一覧, targetRowNum); err != nil {
				return err
			}
		}
	}
	// シートをアクティブに設定
	w.ef.SetActiveSheet(sheetIdx)
	// 上書き保存
	return w.ef.Save()
}
