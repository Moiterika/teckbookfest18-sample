package domain

import (
	"github.com/shopspring/decimal"
)

// Ent仕訳 は仕訳データの1行を表す構造体です
type Ent仕訳 struct {
	FldNo               int64
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
	// ここから追加項目
	*Val仕訳詳細
}

// GetVal仕訳詳細From は引数に指定された仕訳から有効な詳細情報を取得します
func (e *Ent仕訳) GetVal仕訳詳細From(other *Ent仕訳) *Val仕訳詳細 {
	// 勘定科目が以前と異なる場合、無効なので取得しない
	if e.Fld借方勘定科目 != other.Fld借方勘定科目 {
		return nil
	}
	return other.Val仕訳詳細
}

func (e *Ent仕訳) Key() 仕訳Key {
	return 仕訳Key{
		Fld仕訳ID:  e.Fld仕訳ID,
		Fld仕訳行番号: e.Fld仕訳行番号,
	}
}
