package io

import (
	"encoding/csv"
	"fmt"
	"teckbookfest18-sample/domain"
)

// 仕訳CsvReader はQry仕訳インターフェースの実装です
type 仕訳CsvReader struct {
	reader *csv.Reader // CSVリーダーをプライベートフィールドとして保持
}

// New仕訳CsvReader は仕訳CsvReaderの新しいインスタンスを作成します
func New仕訳CsvReader(reader *csv.Reader) domain.Qry仕訳 {
	return &仕訳CsvReader{
		reader: reader,
	}
}

// ReadAll は初期化時に設定されたCSVリーダーから仕訳データを読み取ります
func (q *仕訳CsvReader) ReadAll() ([]*domain.Ent仕訳, error) {
	// CSVリーダーの設定
	q.reader.Comma = ','
	// r.FieldsPerRecord = -1
	q.reader.LazyQuotes = true // RFC 4180 に厳密に従わない不正なクォート（引用符）を無視

	// ヘッダーをスキップ
	if _, err := q.reader.Read(); err != nil {
		return nil, err
	}

	var rows []*csvFreee仕訳
	rowNumber := 2
	for {
		record, err := q.reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("%d行目: %w", rowNumber, err)
		}
		row, err := NewCsvFreee仕訳(record)
		if err != nil {
			return nil, fmt.Errorf("%d行目: %w", rowNumber, err)
		}
		rows = append(rows, row)
		rowNumber++
	}

	// io層のCsvFreee仕訳からdomain層のEnt仕訳に変換
	domainRecords := make([]*domain.Ent仕訳, len(rows))
	for i, row := range rows {
		domainRecords[i] = &domain.Ent仕訳{
			FldNo:       row.FldNo,
			Fld取引日:      row.Fld取引日,
			Fld管理番号:     row.Fld管理番号,
			Fld借方勘定科目:   row.Fld借方勘定科目,
			Fld借方決算書表示名: row.Fld借方決算書表示名,
			Fld借方勘定科目ショートカット1: row.Fld借方勘定科目ショートカット1,
			Fld借方勘定科目ショートカット2: row.Fld借方勘定科目ショートカット2,
			Fld借方金額:             row.Fld借方金額,
			Fld借方税区分:            row.Fld借方税区分,
			Fld借方税金額:            row.Fld借方税金額,
			Fld借方内税外税:           row.Fld借方内税外税,
			Fld借方税率:             row.Fld借方税率,
			Fld借方軽減税率有無:         row.Fld借方軽減税率有無,
			Fld借方取引先コード:         row.Fld借方取引先コード,
			Fld借方取引先名:           row.Fld借方取引先名,
			Fld借方取引先ショートカット1:    row.Fld借方取引先ショートカット1,
			Fld借方取引先ショートカット2:    row.Fld借方取引先ショートカット2,
			Fld借方品目:             row.Fld借方品目,
			Fld借方品目ショートカット1:     row.Fld借方品目ショートカット1,
			Fld借方品目ショートカット2:     row.Fld借方品目ショートカット2,
			Fld借方補助科目名:          row.Fld借方補助科目名,
			Fld借方補助科目ショートカット1:   row.Fld借方補助科目ショートカット1,
			Fld借方補助科目ショートカット2:   row.Fld借方補助科目ショートカット2,
			Fld借方部門:             row.Fld借方部門,
			Fld借方部門ショートカット1:     row.Fld借方部門ショートカット1,
			Fld借方部門ショートカット2:     row.Fld借方部門ショートカット2,
			Fld借方メモ:             row.Fld借方メモ,
			Fld借方メモショートカット1:     row.Fld借方メモショートカット1,
			Fld借方メモショートカット2:     row.Fld借方メモショートカット2,
			Fld借方セグメント1:         row.Fld借方セグメント1,
			Fld借方セグメント1ショートカット1: row.Fld借方セグメント1ショートカット1,
			Fld借方セグメント1ショートカット2: row.Fld借方セグメント1ショートカット2,
			Fld借方セグメント2:         row.Fld借方セグメント2,
			Fld借方セグメント2ショートカット1: row.Fld借方セグメント2ショートカット1,
			Fld借方セグメント2ショートカット2: row.Fld借方セグメント2ショートカット2,
			Fld借方セグメント3:         row.Fld借方セグメント3,
			Fld借方セグメント3ショートカット1: row.Fld借方セグメント3ショートカット1,
			Fld借方セグメント3ショートカット2: row.Fld借方セグメント3ショートカット2,
			Fld借方備考:             row.Fld借方備考,
			Fld貸方勘定科目:           row.Fld貸方勘定科目,
			Fld貸方決算書表示名:         row.Fld貸方決算書表示名,
			Fld貸方勘定科目ショートカット1:   row.Fld貸方勘定科目ショートカット1,
			Fld貸方勘定科目ショートカット2:   row.Fld貸方勘定科目ショートカット2,
			Fld貸方金額:             row.Fld貸方金額,
			Fld貸方税区分:            row.Fld貸方税区分,
			Fld貸方税金額:            row.Fld貸方税金額,
			Fld貸方内税外税:           row.Fld貸方内税外税,
			Fld貸方税率:             row.Fld貸方税率,
			Fld貸方軽減税率有無:         row.Fld貸方軽減税率有無,
			Fld貸方取引先コード:         row.Fld貸方取引先コード,
			Fld貸方取引先名:           row.Fld貸方取引先名,
			Fld貸方取引先ショートカット1:    row.Fld貸方取引先ショートカット1,
			Fld貸方取引先ショートカット2:    row.Fld貸方取引先ショートカット2,
			Fld貸方品目:             row.Fld貸方品目,
			Fld貸方品目ショートカット1:     row.Fld貸方品目ショートカット1,
			Fld貸方品目ショートカット2:     row.Fld貸方品目ショートカット2,
			Fld貸方補助科目名:          row.Fld貸方補助科目名,
			Fld貸方補助科目ショートカット1:   row.Fld貸方補助科目ショートカット1,
			Fld貸方補助科目ショートカット2:   row.Fld貸方補助科目ショートカット2,
			Fld貸方部門:             row.Fld貸方部門,
			Fld貸方部門ショートカット1:     row.Fld貸方部門ショートカット1,
			Fld貸方部門ショートカット2:     row.Fld貸方部門ショートカット2,
			Fld貸方メモ:             row.Fld貸方メモ,
			Fld貸方メモショートカット1:     row.Fld貸方メモショートカット1,
			Fld貸方メモショートカット2:     row.Fld貸方メモショートカット2,
			Fld貸方セグメント1:         row.Fld貸方セグメント1,
			Fld貸方セグメント1ショートカット1: row.Fld貸方セグメント1ショートカット1,
			Fld貸方セグメント1ショートカット2: row.Fld貸方セグメント1ショートカット2,
			Fld貸方セグメント2:         row.Fld貸方セグメント2,
			Fld貸方セグメント2ショートカット1: row.Fld貸方セグメント2ショートカット1,
			Fld貸方セグメント2ショートカット2: row.Fld貸方セグメント2ショートカット2,
			Fld貸方セグメント3:         row.Fld貸方セグメント3,
			Fld貸方セグメント3ショートカット1: row.Fld貸方セグメント3ショートカット1,
			Fld貸方セグメント3ショートカット2: row.Fld貸方セグメント3ショートカット2,
			Fld貸方備考:             row.Fld貸方備考,
			Fld決算整理仕訳:           row.Fld決算整理仕訳,
			Fld発行元:              row.Fld発行元,
			Fld作成日時:             row.Fld作成日時,
			Fld更新日時:             row.Fld更新日時,
			Fld承認状況仕訳承認:         row.Fld承認状況仕訳承認,
			Fld申請者仕訳承認:          row.Fld申請者仕訳承認,
			Fld申請日時仕訳承認:         row.Fld申請日時仕訳承認,
			Fld承認者仕訳承認:          row.Fld承認者仕訳承認,
			Fld承認日時仕訳承認:         row.Fld承認日時仕訳承認,
			Fld作成者:              row.Fld作成者,
			Fld消費税経理処理方法:        row.Fld消費税経理処理方法,
			Fld取引ID:             row.Fld取引ID,
			Fld口座振替ID:           row.Fld口座振替ID,
			Fld振替伝票ID:           row.Fld振替伝票ID,
			Fld仕訳ID:             row.Fld仕訳ID,
			Fld仕訳番号:             row.Fld仕訳番号,
			Fld期末日取引フラグ:         row.Fld期末日取引フラグ,
			Fld取引支払日:            row.Fld取引支払日,
			Fld仕訳行番号:            row.Fld仕訳行番号,
			Fld仕訳行数:             row.Fld仕訳行数,
			Fldレコード番号:           row.Fldレコード番号,
			Fld取引内容:             row.Fld取引内容,
			Fld登録した方法:           row.Fld登録した方法,
			Fld経費精算申請番号:         row.Fld経費精算申請番号,
			Fld支払依頼申請番号:         row.Fld支払依頼申請番号,
		}
	}

	return domainRecords, nil
}
