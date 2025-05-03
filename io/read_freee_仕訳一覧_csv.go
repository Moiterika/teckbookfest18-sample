package io

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// readCsvFreee仕訳一覧 はCSVリーダーから仕訳データを読み取ります
func readCsvFreee仕訳一覧(r *csv.Reader) ([]*CsvFreee仕訳, error) {
	r.Comma = ','
	// r.FieldsPerRecord = -1
	r.LazyQuotes = true // RFC 4180 に厳密に従わない不正なクォート（引用符）を無視

	// ヘッダーをスキップ
	if _, err := r.Read(); err != nil {
		return nil, err
	}

	var records []*CsvFreee仕訳
	rowNumber := 2
	for {
		row, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		借方金額, err := convertToDecimal(row[7])
		if err != nil {
			return nil, fmt.Errorf("%d行目:借方金額が整数以外です。値=%s", rowNumber, row[7])
		}
		借方税金額, err := convertToDecimal(row[9])
		if err != nil {
			return nil, fmt.Errorf("%d行目:借方税金額が整数以外です。値=%s", rowNumber, row[9])
		}
		借方税率, err := convertToDecimal(row[11])
		if err != nil {
			return nil, fmt.Errorf("%d行目:借方税率が整数以外です。値=%s", rowNumber, row[11])
		}
		貸方金額, err := convertToDecimal(row[43])
		if err != nil {
			return nil, fmt.Errorf("%d行目:貸方金額が整数以外です。値=%s", rowNumber, row[43])
		}
		貸方税金額, err := convertToDecimal(row[45])
		if err != nil {
			return nil, fmt.Errorf("%d行目:貸貸方税金額が整数以外です。値=%s", rowNumber, row[45])
		}
		貸方税率, err := convertToDecimal(row[47])
		if err != nil {
			return nil, fmt.Errorf("%d行目:貸方税率が整数以外です。値=%s", rowNumber, row[47])
		}

		rec := CsvFreee仕訳{
			FldNo:       row[0],
			Fld取引日:      row[1],
			Fld管理番号:     row[2],
			Fld借方勘定科目:   row[3],
			Fld借方決算書表示名: row[4],
			Fld借方勘定科目ショートカット1: row[5],
			Fld借方勘定科目ショートカット2: row[6],
			Fld借方金額:             借方金額,
			Fld借方税区分:            row[8],
			Fld借方税金額:            借方税金額,
			Fld借方内税外税:           row[10],
			Fld借方税率:             借方税率,
			Fld借方軽減税率有無:         row[12],
			Fld借方取引先コード:         row[13],
			Fld借方取引先名:           row[14],
			Fld借方取引先ショートカット1:    row[15],
			Fld借方取引先ショートカット2:    row[16],
			Fld借方品目:             row[17],
			Fld借方品目ショートカット1:     row[18],
			Fld借方品目ショートカット2:     row[19],
			Fld借方補助科目名:          row[20],
			Fld借方補助科目ショートカット1:   row[21],
			Fld借方補助科目ショートカット2:   row[22],
			Fld借方部門:             row[23],
			Fld借方部門ショートカット1:     row[24],
			Fld借方部門ショートカット2:     row[25],
			Fld借方メモ:             row[26],
			Fld借方メモショートカット1:     row[27],
			Fld借方メモショートカット2:     row[28],
			Fld借方セグメント1:         row[29],
			Fld借方セグメント1ショートカット1: row[30],
			Fld借方セグメント1ショートカット2: row[31],
			Fld借方セグメント2:         row[32],
			Fld借方セグメント2ショートカット1: row[33],
			Fld借方セグメント2ショートカット2: row[34],
			Fld借方セグメント3:         row[35],
			Fld借方セグメント3ショートカット1: row[36],
			Fld借方セグメント3ショートカット2: row[37],
			Fld借方備考:             row[38],
			Fld貸方勘定科目:           row[39],
			Fld貸方決算書表示名:         row[40],
			Fld貸方勘定科目ショートカット1:   row[41],
			Fld貸方勘定科目ショートカット2:   row[42],
			Fld貸方金額:             貸方金額,
			Fld貸方税区分:            row[44],
			Fld貸方税金額:            貸方税金額,
			Fld貸方内税外税:           row[46],
			Fld貸方税率:             貸方税率,
			Fld貸方軽減税率有無:         row[48],
			Fld貸方取引先コード:         row[49],
			Fld貸方取引先名:           row[50],
			Fld貸方取引先ショートカット1:    row[51],
			Fld貸方取引先ショートカット2:    row[52],
			Fld貸方品目:             row[53],
			Fld貸方品目ショートカット1:     row[54],
			Fld貸方品目ショートカット2:     row[55],
			Fld貸方補助科目名:          row[56],
			Fld貸方補助科目ショートカット1:   row[57],
			Fld貸方補助科目ショートカット2:   row[58],
			Fld貸方部門:             row[59],
			Fld貸方部門ショートカット1:     row[60],
			Fld貸方部門ショートカット2:     row[61],
			Fld貸方メモ:             row[62],
			Fld貸方メモショートカット1:     row[63],
			Fld貸方メモショートカット2:     row[64],
			Fld貸方セグメント1:         row[65],
			Fld貸方セグメント1ショートカット1: row[66],
			Fld貸方セグメント1ショートカット2: row[67],
			Fld貸方セグメント2:         row[68],
			Fld貸方セグメント2ショートカット1: row[69],
			Fld貸方セグメント2ショートカット2: row[70],
			Fld貸方セグメント3:         row[71],
			Fld貸方セグメント3ショートカット1: row[72],
			Fld貸方セグメント3ショートカット2: row[73],
			Fld貸方備考:             row[74],
			Fld決算整理仕訳:           row[75],
			Fld発行元:              row[76],
			Fld作成日時:             row[77],
			Fld更新日時:             row[78],
			Fld承認状況仕訳承認:         row[79],
			Fld申請者仕訳承認:          row[80],
			Fld申請日時仕訳承認:         row[81],
			Fld承認者仕訳承認:          row[82],
			Fld承認日時仕訳承認:         row[83],
			Fld作成者:              row[84],
			Fld消費税経理処理方法:        row[85],
			Fld取引ID:             row[86],
			Fld口座振替ID:           row[87],
			Fld振替伝票ID:           row[88],
			Fld仕訳ID:             row[89],
			Fld仕訳番号:             row[90],
			Fld期末日取引フラグ:         row[91],
			Fld取引支払日:            row[92],
			Fld仕訳行番号:            row[93],
			Fld仕訳行数:             row[94],
			Fldレコード番号:           row[95],
			Fld取引内容:             row[96],
			Fld登録した方法:           row[97],
			Fld経費精算申請番号:         row[98],
			Fld支払依頼申請番号:         row[99],
		}

		records = append(records, &rec)

		rowNumber++
	}
	return records, nil
}

func convertToDecimal(s string) (decimal.Decimal, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		// 空文字なら 0 を設定
		return decimal.NewFromInt(0), nil
	}
	return decimal.NewFromString(s)
}
