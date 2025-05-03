package io

import (
	"encoding/csv"
	"teckbookfest18-sample/domain"
)

// query仕訳 はQuery仕訳インターフェースの実装です
type query仕訳 struct {
	reader *csv.Reader // CSVリーダーをプライベートフィールドとして保持
}

// NewQuery仕訳 はQuery仕訳の新しいインスタンスを作成します
func NewQuery仕訳(reader *csv.Reader) domain.Query仕訳 {
	return &query仕訳{
		reader: reader,
	}
}

// Read は初期化時に設定されたCSVリーダーから仕訳データを読み取ります
func (r *query仕訳) Read() ([]*domain.Ent仕訳, error) {
	// 内部実装で使用するio層の型からdomain層の型に変換
	records, err := readCsvFreee仕訳一覧(r.reader)
	if err != nil {
		return nil, err
	}

	// io層のCsvFreee仕訳からdomain層のEnt仕訳に変換
	domainRecords := make([]*domain.Ent仕訳, len(records))
	for i, record := range records {
		domainRecords[i] = &domain.Ent仕訳{
			FldNo:       record.FldNo,
			Fld取引日:      record.Fld取引日,
			Fld管理番号:     record.Fld管理番号,
			Fld借方勘定科目:   record.Fld借方勘定科目,
			Fld借方決算書表示名: record.Fld借方決算書表示名,
			Fld借方勘定科目ショートカット1: record.Fld借方勘定科目ショートカット1,
			Fld借方勘定科目ショートカット2: record.Fld借方勘定科目ショートカット2,
			Fld借方金額:             record.Fld借方金額,
			Fld借方税区分:            record.Fld借方税区分,
			Fld借方税金額:            record.Fld借方税金額,
			Fld借方内税外税:           record.Fld借方内税外税,
			Fld借方税率:             record.Fld借方税率,
			Fld借方軽減税率有無:         record.Fld借方軽減税率有無,
			Fld借方取引先コード:         record.Fld借方取引先コード,
			Fld借方取引先名:           record.Fld借方取引先名,
			Fld借方取引先ショートカット1:    record.Fld借方取引先ショートカット1,
			Fld借方取引先ショートカット2:    record.Fld借方取引先ショートカット2,
			Fld借方品目:             record.Fld借方品目,
			Fld借方品目ショートカット1:     record.Fld借方品目ショートカット1,
			Fld借方品目ショートカット2:     record.Fld借方品目ショートカット2,
			Fld借方補助科目名:          record.Fld借方補助科目名,
			Fld借方補助科目ショートカット1:   record.Fld借方補助科目ショートカット1,
			Fld借方補助科目ショートカット2:   record.Fld借方補助科目ショートカット2,
			Fld借方部門:             record.Fld借方部門,
			Fld借方部門ショートカット1:     record.Fld借方部門ショートカット1,
			Fld借方部門ショートカット2:     record.Fld借方部門ショートカット2,
			Fld借方メモ:             record.Fld借方メモ,
			Fld借方メモショートカット1:     record.Fld借方メモショートカット1,
			Fld借方メモショートカット2:     record.Fld借方メモショートカット2,
			Fld借方セグメント1:         record.Fld借方セグメント1,
			Fld借方セグメント1ショートカット1: record.Fld借方セグメント1ショートカット1,
			Fld借方セグメント1ショートカット2: record.Fld借方セグメント1ショートカット2,
			Fld借方セグメント2:         record.Fld借方セグメント2,
			Fld借方セグメント2ショートカット1: record.Fld借方セグメント2ショートカット1,
			Fld借方セグメント2ショートカット2: record.Fld借方セグメント2ショートカット2,
			Fld借方セグメント3:         record.Fld借方セグメント3,
			Fld借方セグメント3ショートカット1: record.Fld借方セグメント3ショートカット1,
			Fld借方セグメント3ショートカット2: record.Fld借方セグメント3ショートカット2,
			Fld借方備考:             record.Fld借方備考,
			Fld貸方勘定科目:           record.Fld貸方勘定科目,
			Fld貸方決算書表示名:         record.Fld貸方決算書表示名,
			Fld貸方勘定科目ショートカット1:   record.Fld貸方勘定科目ショートカット1,
			Fld貸方勘定科目ショートカット2:   record.Fld貸方勘定科目ショートカット2,
			Fld貸方金額:             record.Fld貸方金額,
			Fld貸方税区分:            record.Fld貸方税区分,
			Fld貸方税金額:            record.Fld貸方税金額,
			Fld貸方内税外税:           record.Fld貸方内税外税,
			Fld貸方税率:             record.Fld貸方税率,
			Fld貸方軽減税率有無:         record.Fld貸方軽減税率有無,
			Fld貸方取引先コード:         record.Fld貸方取引先コード,
			Fld貸方取引先名:           record.Fld貸方取引先名,
			Fld貸方取引先ショートカット1:    record.Fld貸方取引先ショートカット1,
			Fld貸方取引先ショートカット2:    record.Fld貸方取引先ショートカット2,
			Fld貸方品目:             record.Fld貸方品目,
			Fld貸方品目ショートカット1:     record.Fld貸方品目ショートカット1,
			Fld貸方品目ショートカット2:     record.Fld貸方品目ショートカット2,
			Fld貸方補助科目名:          record.Fld貸方補助科目名,
			Fld貸方補助科目ショートカット1:   record.Fld貸方補助科目ショートカット1,
			Fld貸方補助科目ショートカット2:   record.Fld貸方補助科目ショートカット2,
			Fld貸方部門:             record.Fld貸方部門,
			Fld貸方部門ショートカット1:     record.Fld貸方部門ショートカット1,
			Fld貸方部門ショートカット2:     record.Fld貸方部門ショートカット2,
			Fld貸方メモ:             record.Fld貸方メモ,
			Fld貸方メモショートカット1:     record.Fld貸方メモショートカット1,
			Fld貸方メモショートカット2:     record.Fld貸方メモショートカット2,
			Fld貸方セグメント1:         record.Fld貸方セグメント1,
			Fld貸方セグメント1ショートカット1: record.Fld貸方セグメント1ショートカット1,
			Fld貸方セグメント1ショートカット2: record.Fld貸方セグメント1ショートカット2,
			Fld貸方セグメント2:         record.Fld貸方セグメント2,
			Fld貸方セグメント2ショートカット1: record.Fld貸方セグメント2ショートカット1,
			Fld貸方セグメント2ショートカット2: record.Fld貸方セグメント2ショートカット2,
			Fld貸方セグメント3:         record.Fld貸方セグメント3,
			Fld貸方セグメント3ショートカット1: record.Fld貸方セグメント3ショートカット1,
			Fld貸方セグメント3ショートカット2: record.Fld貸方セグメント3ショートカット2,
			Fld貸方備考:             record.Fld貸方備考,
			Fld決算整理仕訳:           record.Fld決算整理仕訳,
			Fld発行元:              record.Fld発行元,
			Fld作成日時:             record.Fld作成日時,
			Fld更新日時:             record.Fld更新日時,
			Fld承認状況仕訳承認:         record.Fld承認状況仕訳承認,
			Fld申請者仕訳承認:          record.Fld申請者仕訳承認,
			Fld申請日時仕訳承認:         record.Fld申請日時仕訳承認,
			Fld承認者仕訳承認:          record.Fld承認者仕訳承認,
			Fld承認日時仕訳承認:         record.Fld承認日時仕訳承認,
			Fld作成者:              record.Fld作成者,
			Fld消費税経理処理方法:        record.Fld消費税経理処理方法,
			Fld取引ID:             record.Fld取引ID,
			Fld口座振替ID:           record.Fld口座振替ID,
			Fld振替伝票ID:           record.Fld振替伝票ID,
			Fld仕訳ID:             record.Fld仕訳ID,
			Fld仕訳番号:             record.Fld仕訳番号,
			Fld期末日取引フラグ:         record.Fld期末日取引フラグ,
			Fld取引支払日:            record.Fld取引支払日,
			Fld仕訳行番号:            record.Fld仕訳行番号,
			Fld仕訳行数:             record.Fld仕訳行数,
			Fldレコード番号:           record.Fldレコード番号,
			Fld取引内容:             record.Fld取引内容,
			Fld登録した方法:           record.Fld登録した方法,
			Fld経費精算申請番号:         record.Fld経費精算申請番号,
			Fld支払依頼申請番号:         record.Fld支払依頼申請番号,
		}
	}

	return domainRecords, nil
}
