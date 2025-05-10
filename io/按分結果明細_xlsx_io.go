package io

import (
	"fmt"
	"teckbookfest18-sample/domain"

	"github.com/xuri/excelize/v2"
)

// 按分結果明細XlsxIo は I按分結果明細XlsxIo の実装です
// 按分結果明細データのxlsxファイル読み書きを担当します
type 按分結果明細XlsxIo struct {
	ef *excelize.File
}

func New按分結果明細XlsxIo(ef *excelize.File) *按分結果明細XlsxIo {
	return &按分結果明細XlsxIo{ef: ef}
}

const sheet按分結果明細一覧 = "按分結果明細一覧"

// // Read は按分結果明細データを全件読み取ります
// func (x *按分結果明細XlsxIo) Read() ([]*domain.Ent按分結果明細, error) {
// 	rows, err := x.ef.GetRows(sheet按分結果明細一覧)
// 	if err != nil {
// 		return make([]*domain.Ent按分結果明細, 0), err
// 	}
// 	ret := make([]*domain.Ent按分結果明細, 0, len(rows))
// 	for i, row := range rows {
// 		if i == 0 {
// 			// ヘッダーを読み飛ばし
// 			continue
// 		}
// 		var e domain.Ent按分結果明細
// 		e.Fld計上年月 = getStringCell(row, 0)
// 		e.Fld勘定科目 = getStringCell(row, 1)
// 		e.Fldコストプール = getStringCell(row, 2)
// 		e.Fld按分ルール1 = getStringCell(row, 3)
// 		e.Fld按分ルール2 = getStringCell(row, 4)
// 		e.Fld借方税区分 = getStringCell(row, 5)
// 		借方税率, err := getDecimalCell(row, 6)
// 		if err != nil {
// 			return make([]*domain.Ent按分結果明細, 0), fmt.Errorf("%d行目:借方税率エラー: %w", i+1, err)
// 		}
// 		e.Fld借方税率 = 借方税率
// 		合計金額, err := getDecimalCell(row, 7)
// 		if err != nil {
// 			return make([]*domain.Ent按分結果明細, 0), fmt.Errorf("%d行目:合計金額エラー: %w", i+1, err)
// 		}
// 		e.Fld合計金額 = 合計金額
// 		e.Fld按分先 = getStringCell(row, 8)
// 		按分基準値, err := getDecimalCell(row, 9)
// 		if err != nil {
// 			return make([]*domain.Ent按分結果明細, 0), fmt.Errorf("%d行目:按分基準値エラー: %w", i+1, err)
// 		}
// 		e.Fld按分基準値 = 按分基準値
// 		按分誤差, err := getDecimalCell(row, 10)
// 		if err != nil {
// 			return make([]*domain.Ent按分結果明細, 0), fmt.Errorf("%d行目:按分誤差エラー: %w", i+1, err)
// 		}
// 		e.Fld按分誤差 = 按分誤差
// 		按分結果, err := getDecimalCell(row, 11)
// 		if err != nil {
// 			return make([]*domain.Ent按分結果明細, 0), fmt.Errorf("%d行目:按分結果エラー: %w", i+1, err)
// 		}
// 		e.Fld按分結果 = 按分結果
// 		ret = append(ret, &e)
// 	}
// 	return ret, nil
// }

// Save は按分結果明細データを保存します
func (x *按分結果明細XlsxIo) Save(明細一覧 []*domain.Ent按分結果明細) error {
	var sheetIdx int
	var existingRowsCount int
	if idx, err := x.ef.GetSheetIndex(sheet按分結果明細一覧); err == nil {
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
	headers := []interface{}{"計上年月", "勘定科目", "コストプール", "按分ルール1", "按分ルール2", "借方税区分", "借方税率", "合計金額", "按分先", "按分基準値", "按分誤差", "按分結果"}
	x.ef.SetSheetRow(sheet按分結果明細一覧, "A1", &headers)

	// データ行を書き込み
	for i, e := range 明細一覧 {
		row := []interface{}{
			e.Fld計上年月, e.Fld勘定科目, e.Fldコストプール, e.Fld按分ルール1, e.Fld按分ルール2, e.Fld借方税区分, e.Fld借方税率.String(), e.Fld合計金額.String(), e.Fld按分先, e.Fld按分基準値.String(), e.Fld按分誤差.String(), e.Fld按分結果.String(),
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
