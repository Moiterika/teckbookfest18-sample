package io

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// csvFreee仕訳 は仕訳データCSVの行を表す構造体です
type csvFreee仕訳 struct {
	FldNo               string
	Fld取引日              string
	Fld管理番号             string
	Fld借方勘定科目           string
	Fld借方決算書表示名         string
	Fld借方勘定科目ショートカット1   string
	Fld借方勘定科目ショートカット2   string
	Fld借方金額             decimal.Decimal
	Fld借方税区分            string
	Fld借方税金額            decimal.Decimal
	Fld借方内税外税           string
	Fld借方税率             decimal.Decimal
	Fld借方軽減税率有無         string
	Fld借方取引先コード         string
	Fld借方取引先名           string
	Fld借方取引先ショートカット1    string
	Fld借方取引先ショートカット2    string
	Fld借方品目             string
	Fld借方品目ショートカット1     string
	Fld借方品目ショートカット2     string
	Fld借方補助科目名          string
	Fld借方補助科目ショートカット1   string
	Fld借方補助科目ショートカット2   string
	Fld借方部門             string
	Fld借方部門ショートカット1     string
	Fld借方部門ショートカット2     string
	Fld借方メモ             string
	Fld借方メモショートカット1     string
	Fld借方メモショートカット2     string
	Fld借方セグメント1         string
	Fld借方セグメント1ショートカット1 string
	Fld借方セグメント1ショートカット2 string
	Fld借方セグメント2         string
	Fld借方セグメント2ショートカット1 string
	Fld借方セグメント2ショートカット2 string
	Fld借方セグメント3         string
	Fld借方セグメント3ショートカット1 string
	Fld借方セグメント3ショートカット2 string
	Fld借方備考             string
	Fld貸方勘定科目           string
	Fld貸方決算書表示名         string
	Fld貸方勘定科目ショートカット1   string
	Fld貸方勘定科目ショートカット2   string
	Fld貸方金額             decimal.Decimal
	Fld貸方税区分            string
	Fld貸方税金額            decimal.Decimal
	Fld貸方内税外税           string
	Fld貸方税率             decimal.Decimal
	Fld貸方軽減税率有無         string
	Fld貸方取引先コード         string
	Fld貸方取引先名           string
	Fld貸方取引先ショートカット1    string
	Fld貸方取引先ショートカット2    string
	Fld貸方品目             string
	Fld貸方品目ショートカット1     string
	Fld貸方品目ショートカット2     string
	Fld貸方補助科目名          string
	Fld貸方補助科目ショートカット1   string
	Fld貸方補助科目ショートカット2   string
	Fld貸方部門             string
	Fld貸方部門ショートカット1     string
	Fld貸方部門ショートカット2     string
	Fld貸方メモ             string
	Fld貸方メモショートカット1     string
	Fld貸方メモショートカット2     string
	Fld貸方セグメント1         string
	Fld貸方セグメント1ショートカット1 string
	Fld貸方セグメント1ショートカット2 string
	Fld貸方セグメント2         string
	Fld貸方セグメント2ショートカット1 string
	Fld貸方セグメント2ショートカット2 string
	Fld貸方セグメント3         string
	Fld貸方セグメント3ショートカット1 string
	Fld貸方セグメント3ショートカット2 string
	Fld貸方備考             string
	Fld決算整理仕訳           string
	Fld発行元              string
	Fld作成日時             string
	Fld更新日時             string
	Fld承認状況仕訳承認         string
	Fld申請者仕訳承認          string
	Fld申請日時仕訳承認         string
	Fld承認者仕訳承認          string
	Fld承認日時仕訳承認         string
	Fld作成者              string
	Fld消費税経理処理方法        string
	Fld取引ID             string
	Fld口座振替ID           string
	Fld振替伝票ID           string
	Fld仕訳ID             string
	Fld仕訳番号             string
	Fld期末日取引フラグ         string
	Fld取引支払日            string
	Fld仕訳行番号            string
	Fld仕訳行数             string
	Fldレコード番号           string
	Fld取引内容             string
	Fld登録した方法           string
	Fld経費精算申請番号         string
	Fld支払依頼申請番号         string
}

func NewCsvFreee仕訳(row []string) (*csvFreee仕訳, error) {
	借方金額, err := convertToDecimal(row[7])
	if err != nil {
		return nil, fmt.Errorf("借方金額が整数以外です。値=%s", row[7])
	}
	借方税金額, err := convertToDecimal(row[9])
	if err != nil {
		return nil, fmt.Errorf("借方税金額が整数以外です。値=%s", row[9])
	}
	借方税率, err := convertToDecimal(row[11])
	if err != nil {
		return nil, fmt.Errorf("借方税率が整数以外です。値=%s", row[11])
	}
	貸方金額, err := convertToDecimal(row[43])
	if err != nil {
		return nil, fmt.Errorf("貸方金額が整数以外です。値=%s", row[43])
	}
	貸方税金額, err := convertToDecimal(row[45])
	if err != nil {
		return nil, fmt.Errorf("貸貸方税金額が整数以外です。値=%s", row[45])
	}
	貸方税率, err := convertToDecimal(row[47])
	if err != nil {
		return nil, fmt.Errorf("貸方税率が整数以外です。値=%s", row[47])
	}

	ret := csvFreee仕訳{
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
	return &ret, nil
}
