package io

import (
	"fmt"
	"teckbookfest18-sample/domain"

	"github.com/xuri/excelize/v2"
)

// 仕訳XlsxIo は Rep仕訳 の実装です
// 仕訳データのxlsxファイル読み書きを担当します
type 仕訳XlsxIo struct {
	ef *excelize.File
}

func New仕訳XlsxIo(ef *excelize.File) *仕訳XlsxIo {
	return &仕訳XlsxIo{ef: ef}
}

const sheet仕訳一覧 = "仕訳一覧"

// Read仕訳一覧 は仕訳データを全件読み取ります
func (x *仕訳XlsxIo) Read仕訳一覧() ([]*domain.Ent仕訳, error) {
	rows, err := x.ef.GetRows(sheet仕訳一覧)
	if err != nil {
		return make([]*domain.Ent仕訳, 0), err
	}
	ret := make([]*domain.Ent仕訳, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			// ヘッダーを読み飛ばし
			continue
		}
		rowNumber := i + 1
		no, err := getInt64Cell(row, FldIdxNo)
		if err != nil {
			return make([]*domain.Ent仕訳, 0), fmt.Errorf("%d行目:Noエラー: %w", rowNumber, err)
		}
		if no == 0 {
			continue
		}
		借方金額, err := getDecimalCell(row, FldIdx借方金額)
		if err != nil {
			return make([]*domain.Ent仕訳, 0), fmt.Errorf("%d行目:借方金額エラー: %w", rowNumber, err)
		}
		借方税金額, err := getDecimalCell(row, FldIdx借方税金額)
		if err != nil {
			return make([]*domain.Ent仕訳, 0), fmt.Errorf("%d行目:借方税金額エラー: %w", rowNumber, err)
		}
		借方税率, err := getDecimalCell(row, FldIdx借方税率)
		if err != nil {
			return make([]*domain.Ent仕訳, 0), fmt.Errorf("%d行目:借方税率エラー: %w", rowNumber, err)
		}
		貸方金額, err := getDecimalCell(row, FldIdx貸方金額)
		if err != nil {
			return make([]*domain.Ent仕訳, 0), fmt.Errorf("%d行目:貸方金額エラー: %w", rowNumber, err)
		}
		貸方税金額, err := getDecimalCell(row, FldIdx貸方税金額)
		if err != nil {
			return make([]*domain.Ent仕訳, 0), fmt.Errorf("%d行目:貸方税金額エラー: %w", rowNumber, err)
		}
		貸方税率, err := getDecimalCell(row, FldIdx貸方税率)
		if err != nil {
			return make([]*domain.Ent仕訳, 0), fmt.Errorf("%d行目:貸方税率エラー: %w", rowNumber, err)
		}

		var e domain.Ent仕訳
		// 仕訳詳細のフィールドをセット
		e.Val仕訳詳細 = newVal仕訳詳細(getStringCell(row, FldIdx計上年月), getStringCell(row, FldIdx原価要素), getStringCell(row, FldIdxコストプール), getStringCell(row, FldIdx按分ルール1), getStringCell(row, FldIdx按分ルール2))
		e.FldNo = no
		e.Fld取引日 = getStringCell(row, FldIdx取引日)
		e.Fld管理番号 = getStringCell(row, FldIdx管理番号)
		e.Fld借方勘定科目 = getStringCell(row, FldIdx借方勘定科目)
		e.Fld借方決算書表示名 = getStringCell(row, FldIdx借方決算書表示名)
		e.Fld借方勘定科目ショートカット1 = getStringCell(row, FldIdx借方勘定科目ショートカット1)
		e.Fld借方勘定科目ショートカット2 = getStringCell(row, FldIdx借方勘定科目ショートカット2)
		e.Fld借方金額 = 借方金額
		e.Fld借方税区分 = getStringCell(row, FldIdx借方税区分)
		e.Fld借方税金額 = 借方税金額
		e.Fld借方内税外税 = getStringCell(row, FldIdx借方内税外税)
		e.Fld借方税率 = 借方税率
		e.Fld借方軽減税率有無 = getStringCell(row, FldIdx借方軽減税率有無)
		e.Fld借方取引先コード = getStringCell(row, FldIdx借方取引先コード)
		e.Fld借方取引先名 = getStringCell(row, FldIdx借方取引先名)
		e.Fld借方取引先ショートカット1 = getStringCell(row, FldIdx借方取引先ショートカット1)
		e.Fld借方取引先ショートカット2 = getStringCell(row, FldIdx借方取引先ショートカット2)
		e.Fld借方品目 = getStringCell(row, FldIdx借方品目)
		e.Fld借方品目ショートカット1 = getStringCell(row, FldIdx借方品目ショートカット1)
		e.Fld借方品目ショートカット2 = getStringCell(row, FldIdx借方品目ショートカット2)
		e.Fld借方補助科目名 = getStringCell(row, FldIdx借方補助科目名)
		e.Fld借方補助科目ショートカット1 = getStringCell(row, FldIdx借方補助科目ショートカット1)
		e.Fld借方補助科目ショートカット2 = getStringCell(row, FldIdx借方補助科目ショートカット2)
		e.Fld借方部門 = getStringCell(row, FldIdx借方部門)
		e.Fld借方部門ショートカット1 = getStringCell(row, FldIdx借方部門ショートカット1)
		e.Fld借方部門ショートカット2 = getStringCell(row, FldIdx借方部門ショートカット2)
		e.Fld借方メモ = getStringCell(row, FldIdx借方メモ)
		e.Fld借方メモショートカット1 = getStringCell(row, FldIdx借方メモショートカット1)
		e.Fld借方メモショートカット2 = getStringCell(row, FldIdx借方メモショートカット2)
		e.Fld借方セグメント1 = getStringCell(row, FldIdx借方セグメント1)
		e.Fld借方セグメント1ショートカット1 = getStringCell(row, FldIdx借方セグメント1ショートカット1)
		e.Fld借方セグメント1ショートカット2 = getStringCell(row, FldIdx借方セグメント1ショートカット2)
		e.Fld借方セグメント2 = getStringCell(row, FldIdx借方セグメント2)
		e.Fld借方セグメント2ショートカット1 = getStringCell(row, FldIdx借方セグメント2ショートカット1)
		e.Fld借方セグメント2ショートカット2 = getStringCell(row, FldIdx借方セグメント2ショートカット2)
		e.Fld借方セグメント3 = getStringCell(row, FldIdx借方セグメント3)
		e.Fld借方セグメント3ショートカット1 = getStringCell(row, FldIdx借方セグメント3ショートカット1)
		e.Fld借方セグメント3ショートカット2 = getStringCell(row, FldIdx借方セグメント3ショートカット2)
		e.Fld借方備考 = getStringCell(row, FldIdx借方備考)
		e.Fld貸方勘定科目 = getStringCell(row, FldIdx貸方勘定科目)
		e.Fld貸方決算書表示名 = getStringCell(row, FldIdx貸方決算書表示名)
		e.Fld貸方勘定科目ショートカット1 = getStringCell(row, FldIdx貸方勘定科目ショートカット1)
		e.Fld貸方勘定科目ショートカット2 = getStringCell(row, FldIdx貸方勘定科目ショートカット2)
		e.Fld貸方金額 = 貸方金額
		e.Fld貸方税区分 = getStringCell(row, FldIdx貸方税区分)
		e.Fld貸方税金額 = 貸方税金額
		e.Fld貸方内税外税 = getStringCell(row, FldIdx貸方内税外税)
		e.Fld貸方税率 = 貸方税率
		e.Fld貸方軽減税率有無 = getStringCell(row, FldIdx貸方軽減税率有無)
		e.Fld貸方取引先コード = getStringCell(row, FldIdx貸方取引先コード)
		e.Fld貸方取引先名 = getStringCell(row, FldIdx貸方取引先名)
		e.Fld貸方取引先ショートカット1 = getStringCell(row, FldIdx貸方取引先ショートカット1)
		e.Fld貸方取引先ショートカット2 = getStringCell(row, FldIdx貸方取引先ショートカット2)
		e.Fld貸方品目 = getStringCell(row, FldIdx貸方品目)
		e.Fld貸方品目ショートカット1 = getStringCell(row, FldIdx貸方品目ショートカット1)
		e.Fld貸方品目ショートカット2 = getStringCell(row, FldIdx貸方品目ショートカット2)
		e.Fld貸方補助科目名 = getStringCell(row, FldIdx貸方補助科目名)
		e.Fld貸方補助科目ショートカット1 = getStringCell(row, FldIdx貸方補助科目ショートカット1)
		e.Fld貸方補助科目ショートカット2 = getStringCell(row, FldIdx貸方補助科目ショートカット2)
		e.Fld貸方部門 = getStringCell(row, FldIdx貸方部門)
		e.Fld貸方部門ショートカット1 = getStringCell(row, FldIdx貸方部門ショートカット1)
		e.Fld貸方部門ショートカット2 = getStringCell(row, FldIdx貸方部門ショートカット2)
		e.Fld貸方メモ = getStringCell(row, FldIdx貸方メモ)
		e.Fld貸方メモショートカット1 = getStringCell(row, FldIdx貸方メモショートカット1)
		e.Fld貸方メモショートカット2 = getStringCell(row, FldIdx貸方メモショートカット2)
		e.Fld貸方セグメント1 = getStringCell(row, FldIdx貸方セグメント1)
		e.Fld貸方セグメント1ショートカット1 = getStringCell(row, FldIdx貸方セグメント1ショートカット1)
		e.Fld貸方セグメント1ショートカット2 = getStringCell(row, FldIdx貸方セグメント1ショートカット2)
		e.Fld貸方セグメント2 = getStringCell(row, FldIdx貸方セグメント2)
		e.Fld貸方セグメント2ショートカット1 = getStringCell(row, FldIdx貸方セグメント2ショートカット1)
		e.Fld貸方セグメント2ショートカット2 = getStringCell(row, FldIdx貸方セグメント2ショートカット2)
		e.Fld貸方セグメント3 = getStringCell(row, FldIdx貸方セグメント3)
		e.Fld貸方セグメント3ショートカット1 = getStringCell(row, FldIdx貸方セグメント3ショートカット1)
		e.Fld貸方セグメント3ショートカット2 = getStringCell(row, FldIdx貸方セグメント3ショートカット2)
		e.Fld貸方備考 = getStringCell(row, FldIdx貸方備考)
		e.Fld決算整理仕訳 = getStringCell(row, FldIdx決算整理仕訳)
		e.Fld発行元 = getStringCell(row, FldIdx発行元)
		e.Fld作成日時 = getStringCell(row, FldIdx作成日時)
		e.Fld更新日時 = getStringCell(row, FldIdx更新日時)
		e.Fld承認状況仕訳承認 = getStringCell(row, FldIdx承認状況仕訳承認)
		e.Fld申請者仕訳承認 = getStringCell(row, FldIdx申請者仕訳承認)
		e.Fld申請日時仕訳承認 = getStringCell(row, FldIdx申請日時仕訳承認)
		e.Fld承認者仕訳承認 = getStringCell(row, FldIdx承認者仕訳承認)
		e.Fld承認日時仕訳承認 = getStringCell(row, FldIdx承認日時仕訳承認)
		e.Fld作成者 = getStringCell(row, FldIdx作成者)
		e.Fld消費税経理処理方法 = getStringCell(row, FldIdx消費税経理処理方法)
		e.Fld取引ID = getStringCell(row, FldIdx取引ID)
		e.Fld口座振替ID = getStringCell(row, FldIdx口座振替ID)
		e.Fld振替伝票ID = getStringCell(row, FldIdx振替伝票ID)
		e.Fld仕訳ID = getStringCell(row, FldIdx仕訳ID)
		e.Fld仕訳番号 = getStringCell(row, FldIdx仕訳番号)
		e.Fld期末日取引フラグ = getStringCell(row, FldIdx期末日取引フラグ)
		e.Fld取引支払日 = getStringCell(row, FldIdx取引支払日)
		e.Fld仕訳行番号 = getStringCell(row, FldIdx仕訳行番号)
		e.Fld仕訳行数 = getStringCell(row, FldIdx仕訳行数)
		e.Fldレコード番号 = getStringCell(row, FldIdxレコード番号)
		e.Fld取引内容 = getStringCell(row, FldIdx取引内容)
		e.Fld登録した方法 = getStringCell(row, FldIdx登録した方法)
		e.Fld経費精算申請番号 = getStringCell(row, FldIdx経費精算申請番号)
		e.Fld支払依頼申請番号 = getStringCell(row, FldIdx支払依頼申請番号)
		ret = append(ret, &e)
	}
	return ret, nil
}

// Save は仕訳データを保存します
func (x *仕訳XlsxIo) Save(仕訳一覧 []*domain.Ent仕訳) error {
	// シートの有無チェックと既存行取得
	var sheetIdx int
	var existingRowsCount int
	if idx, err := x.ef.GetSheetIndex(sheet仕訳一覧); err != nil {
		return err
	} else if idx != -1 {
		sheetIdx = idx
		existingRows, _ := x.ef.GetRows(sheet仕訳一覧)
		existingRowsCount = len(existingRows)
	} else {
		var err error
		sheetIdx, err = x.ef.NewSheet(sheet仕訳一覧)
		if err != nil {
			return err
		}
	}
	// ヘッダー行を書き込み
	headers := []interface{}{
		"計上年月", "原価要素", "コストプール", "按分ルール1", "按分ルール2", "No", "取引日", "管理番号",
		"借方勘定科目", "借方決算書表示名", "借方勘定科目ショートカット1", "借方勘定科目ショートカット2",
		"借方金額", "借方税区分", "借方税金額", "借方内税外税", "借方税率", "借方軽減税率有無",
		"借方取引先コード", "借方取引先名", "借方取引先ショートカット1", "借方取引先ショートカット2",
		"借方品目", "借方品目ショートカット1", "借方品目ショートカット2",
		"借方補助科目名", "借方補助科目ショートカット1", "借方補助科目ショートカット2",
		"借方部門", "借方部門ショートカット1", "借方部門ショートカット2",
		"借方メモ", "借方メモショートカット1", "借方メモショートカット2",
		"借方セグメント1", "借方セグメント1ショートカット1", "借方セグメント1ショートカット2",
		"借方セグメント2", "借方セグメント2ショートカット1", "借方セグメント2ショートカット2",
		"借方セグメント3", "借方セグメント3ショートカット1", "借方セグメント3ショートカット2",
		"借方備考",
		"貸方勘定科目", "貸方決算書表示名", "貸方勘定科目ショートカット1", "貸方勘定科目ショートカット2",
		"貸方金額", "貸方税区分", "貸方税金額", "貸方内税外税", "貸方税率", "貸方軽減税率有無",
		"貸方取引先コード", "貸方取引先名", "貸方取引先ショートカット1", "貸方取引先ショートカット2",
		"貸方品目", "貸方品目ショートカット1", "貸方品目ショートカット2",
		"貸方補助科目名", "貸方補助科目ショートカット1", "貸方補助科目ショートカット2",
		"貸方部門", "貸方部門ショートカット1", "貸方部門ショートカット2",
		"貸方メモ", "貸方メモショートカット1", "貸方メモショートカット2",
		"貸方セグメント1", "貸方セグメント1ショートカット1", "貸方セグメント1ショートカット2",
		"貸方セグメント2", "貸方セグメント2ショートカット1", "貸方セグメント2ショートカット2",
		"貸方セグメント3", "貸方セグメント3ショートカット1", "貸方セグメント3ショートカット2",
		"貸方備考",
		"決算整理仕訳", "発行元", "作成日時", "更新日時",
		"承認状況仕訳承認", "申請者仕訳承認", "申請日時仕訳承認", "承認者仕訳承認", "承認日時仕訳承認",
		"作成者", "消費税経理処理方法",
		"取引ID", "口座振替ID", "振替伝票ID", "仕訳ID", "仕訳番号", "期末日取引フラグ", "取引支払日",
		"仕訳行番号", "仕訳行数", "レコード番号", "取引内容", "登録した方法", "経費精算申請番号", "支払依頼申請番号",
	}
	x.ef.SetSheetRow(sheet仕訳一覧, "A1", &headers)

	// データ行を書き込み
	for i, e := range 仕訳一覧 {
		// Val仕訳詳細がnilの場合は初期値で埋める
		var d *domain.Val仕訳詳細
		if e.Val仕訳詳細 == nil {
			// 日本語コメント: Val仕訳詳細がnilの場合は空の構造体で初期化
			d = &domain.Val仕訳詳細{}
		} else {
			d = e.Val仕訳詳細
		}
		row := []interface{}{
			d.Fld計上年月, d.Fld原価要素, d.Fldコストプール, d.Fld按分ルール1, d.Fld按分ルール2,
			e.FldNo, e.Fld取引日, e.Fld管理番号,
			e.Fld借方勘定科目, e.Fld借方決算書表示名, e.Fld借方勘定科目ショートカット1, e.Fld借方勘定科目ショートカット2,
			e.Fld借方金額.IntPart(), e.Fld借方税区分, e.Fld借方税金額.IntPart(), e.Fld借方内税外税, e.Fld借方税率, e.Fld借方軽減税率有無,
			e.Fld借方取引先コード, e.Fld借方取引先名, e.Fld借方取引先ショートカット1, e.Fld借方取引先ショートカット2,
			e.Fld借方品目, e.Fld借方品目ショートカット1, e.Fld借方品目ショートカット2,
			e.Fld借方補助科目名, e.Fld借方補助科目ショートカット1, e.Fld借方補助科目ショートカット2,
			e.Fld借方部門, e.Fld借方部門ショートカット1, e.Fld借方部門ショートカット2,
			e.Fld借方メモ, e.Fld借方メモショートカット1, e.Fld借方メモショートカット2,
			e.Fld借方セグメント1, e.Fld借方セグメント1ショートカット1, e.Fld借方セグメント1ショートカット2,
			e.Fld借方セグメント2, e.Fld借方セグメント2ショートカット1, e.Fld借方セグメント2ショートカット2,
			e.Fld借方セグメント3, e.Fld借方セグメント3ショートカット1, e.Fld借方セグメント3ショートカット2,
			e.Fld借方備考,
			e.Fld貸方勘定科目, e.Fld貸方決算書表示名, e.Fld貸方勘定科目ショートカット1, e.Fld貸方勘定科目ショートカット2,
			e.Fld貸方金額.IntPart(), e.Fld貸方税区分, e.Fld貸方税金額.IntPart(), e.Fld貸方内税外税, e.Fld貸方税率, e.Fld貸方軽減税率有無,
			e.Fld貸方取引先コード, e.Fld貸方取引先名, e.Fld貸方取引先ショートカット1, e.Fld貸方取引先ショートカット2,
			e.Fld貸方品目, e.Fld貸方品目ショートカット1, e.Fld貸方品目ショートカット2,
			e.Fld貸方補助科目名, e.Fld貸方補助科目ショートカット1, e.Fld貸方補助科目ショートカット2,
			e.Fld貸方部門, e.Fld貸方部門ショートカット1, e.Fld貸方部門ショートカット2,
			e.Fld貸方メモ, e.Fld貸方メモショートカット1, e.Fld貸方メモショートカット2,
			e.Fld貸方セグメント1, e.Fld貸方セグメント1ショートカット1, e.Fld貸方セグメント1ショートカット2,
			e.Fld貸方セグメント2, e.Fld貸方セグメント2ショートカット1, e.Fld貸方セグメント2ショートカット2,
			e.Fld貸方セグメント3, e.Fld貸方セグメント3ショートカット1, e.Fld貸方セグメント3ショートカット2,
			e.Fld貸方備考,
			e.Fld決算整理仕訳, e.Fld発行元, e.Fld作成日時, e.Fld更新日時,
			e.Fld承認状況仕訳承認, e.Fld申請者仕訳承認, e.Fld申請日時仕訳承認, e.Fld承認者仕訳承認, e.Fld承認日時仕訳承認,
			e.Fld作成者, e.Fld消費税経理処理方法,
			e.Fld取引ID, e.Fld口座振替ID, e.Fld振替伝票ID, e.Fld仕訳ID, e.Fld仕訳番号, e.Fld期末日取引フラグ, e.Fld取引支払日,
			e.Fld仕訳行番号, e.Fld仕訳行数, e.Fldレコード番号, e.Fld取引内容, e.Fld登録した方法, e.Fld経費精算申請番号, e.Fld支払依頼申請番号,
		}
		cell := fmt.Sprintf("A%d", i+2)
		x.ef.SetSheetRow(sheet仕訳一覧, cell, &row)
	}

	// 余分な行を削除（データ数が減った場合）
	if existingRowsCount > len(仕訳一覧)+1 {
		targetRow := len(仕訳一覧) + 2
		count := existingRowsCount - (len(仕訳一覧) + 1)
		for i := 0; i < count; i++ {
			if err := x.ef.RemoveRow(sheet仕訳一覧, targetRow); err != nil {
				return err
			}
		}
	}
	// シートをアクティブに設定
	x.ef.SetActiveSheet(sheetIdx)
	// 上書き保存
	return x.ef.Save()
}
