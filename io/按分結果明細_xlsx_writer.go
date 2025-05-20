package io

import (
	"fmt"
	"teckbookfest18-sample/domain"

	"github.com/xuri/excelize/v2"
)

// 按分結果明細XlsxWriter はCmd按分結果明細の実装です
// 按分結果明細データのxlsxファイルへの保存のみを担当します
// 読み込み処理は実装しません
// 実体はio層で管理し、domain層のインターフェースに従います
type 按分結果明細XlsxWriter struct {
	// excelize.Fileの参照を保持
	ef *excelize.File
}

// New按分結果明細XlsxWriterは按分結果明細XlsxWriterのコンストラクタです
func New按分結果明細XlsxWriter(ef *excelize.File) *按分結果明細XlsxWriter {
	return &按分結果明細XlsxWriter{ef: ef}
}

const sheet按分結果明細一覧 = "按分結果明細一覧"

// Saveは按分結果明細データをxlsxファイルに保存します
// ヘッダー行・データ行を書き込み、余分な行は削除します
// 保存後、シートをアクティブに設定しファイルを上書き保存します
func (x *按分結果明細XlsxWriter) Save(明細一覧 []*domain.Ent按分結果明細) error {
	var sheetIdx int
	var existingRowsCount int
	if idx, err := x.ef.GetSheetIndex(sheet按分結果明細一覧); err != nil {
		return err
	} else if idx != -1 {
		sheetIdx = idx
		existingRows, _ := x.ef.GetRows(sheet按分結果明細一覧)
		existingRowsCount = len(existingRows)
	} else {
		var err error
		sheetIdx, err = x.ef.NewSheet(sheet按分結果明細一覧)
		if err != nil {
			return err
		}
	}
	// ヘッダー行を書き込み
	headers := []interface{}{"計上年月", "原価要素", "コストプール", "按分ルール1", "按分ルール2", "直間", "借方税区分", "借方税率", "合計金額", "按分先", "按分基準値", "按分誤差", "按分結果"}
	x.ef.SetSheetRow(sheet按分結果明細一覧, "A1", &headers)

	// データ行を書き込み
	for i, e := range 明細一覧 {
		var 直間 string
		if e.FldIs直接費 {
			直間 = "直接費"
		} else {
			直間 = "間接費"
		}
		row := []interface{}{
			e.Fld計上年月, e.Fld原価要素, e.Fldコストプール, e.Fld按分ルール1, e.Fld按分ルール2, 直間, e.Fld借方税区分, e.Fld借方税率.String(), e.Fld合計金額.IntPart(), e.Fld按分先, e.Fld按分基準値.String(), e.Fld按分誤差.IntPart(), e.Fld按分結果.IntPart(),
		}
		cell := fmt.Sprintf("A%d", i+2)
		x.ef.SetSheetRow(sheet按分結果明細一覧, cell, &row)
	}

	// 余分な行を削除（データ数が減った場合）
	if existingRowsCount > len(明細一覧)+1 {
		targetRow := len(明細一覧) + 2
		count := existingRowsCount - (len(明細一覧) + 1)
		for i := 0; i < count; i++ {
			if err := x.ef.RemoveRow(sheet按分結果明細一覧, targetRow); err != nil {
				return err
			}
		}
	}
	// シートをアクティブに設定
	x.ef.SetActiveSheet(sheetIdx)
	// 上書き保存
	return x.ef.Save()
}
