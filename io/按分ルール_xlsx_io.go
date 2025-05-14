package io

import (
	"fmt"
	"teckbookfest18-sample/domain"

	"github.com/xuri/excelize/v2"
)

// 按分ルールXlsxIo は I按分ルールXlsxIo の実装です
// 按分ルールデータのxlsxファイル読み書きを担当します
type 按分ルールXlsxIo struct {
	ef *excelize.File
}

func New按分ルールXlsxIo(ef *excelize.File) *按分ルールXlsxIo {
	return &按分ルールXlsxIo{ef: ef}
}

const sheet按分ルール一覧 = "按分ルール一覧"

// Read按分ルール一覧 は按分ルールデータを全件読み取ります
func (x *按分ルールXlsxIo) Read按分ルール一覧() ([]*domain.Ent按分ルール, error) {
	rows, err := x.ef.GetRows(sheet按分ルール一覧)
	if err != nil {
		return make([]*domain.Ent按分ルール, 0), err
	}
	ret := make([]*domain.Ent按分ルール, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			// ヘッダーを読み飛ばし
			continue
		}
		var e domain.Ent按分ルール
		e.Fld按分ルール1 = getStringCell(row, 0)
		e.Fld按分ルール2 = getStringCell(row, 1)
		e.Fld按分先 = getStringCell(row, 2)
		按分基準値, err := getDecimalCell(row, 3)
		if err != nil {
			return make([]*domain.Ent按分ルール, 0), fmt.Errorf("%d行目:値エラー: %w", i+1, err)
		}
		e.Fld按分基準値 = 按分基準値
		ret = append(ret, &e)
	}
	return ret, nil
}

// Save は按分ルールデータを保存します
func (x *按分ルールXlsxIo) Save(按分ルール一覧 []*domain.Ent按分ルール) error {
	// シートの有無チェックと既存行取得
	var sheetIdx int
	var existingRowsCount int
	if idx, err := x.ef.GetSheetIndex(sheet按分ルール一覧); err != nil {
		return err
	} else if idx != -1 {
		sheetIdx = idx
		existingRows, _ := x.ef.GetRows(sheet按分ルール一覧)
		existingRowsCount = len(existingRows)
	} else {
		var err error
		sheetIdx, err = x.ef.NewSheet(sheet按分ルール一覧)
		if err != nil {
			return err
		}
	}
	// ヘッダー行を書き込み
	headers := []interface{}{"按分ルール1", "按分ルール2", "按分先", "按分基準値"}
	x.ef.SetSheetRow(sheet按分ルール一覧, "A1", &headers)

	// データ行を書き込み
	for i, e := range 按分ルール一覧 {
		row := []interface{}{
			e.Fld按分ルール1, e.Fld按分ルール2, e.Fld按分先, e.Fld按分基準値.IntPart(),
		}
		cell := fmt.Sprintf("A%d", i+2)
		x.ef.SetSheetRow(sheet按分ルール一覧, cell, &row)
	}

	// 余分な行を削除（データ数が減った場合）
	if existingRowsCount > len(按分ルール一覧)+1 {
		targetRow := len(按分ルール一覧) + 2
		count := existingRowsCount - (len(按分ルール一覧) + 1)
		for i := 0; i < count; i++ {
			if err := x.ef.RemoveRow(sheet按分ルール一覧, targetRow); err != nil {
				return err
			}
		}
	}
	// シートをアクティブに設定
	x.ef.SetActiveSheet(sheetIdx)
	// 上書き保存
	return x.ef.Save()
}
