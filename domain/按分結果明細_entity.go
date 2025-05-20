package domain

import "github.com/shopspring/decimal"

type Ent按分結果明細 struct {
	Fld計上年月   string
	Fld原価要素   string
	Fldコストプール string
	Fld按分ルール1 string
	Fld按分ルール2 string
	FldIs直接費  bool
	Fld借方税区分  string
	Fld借方税率   decimal.Decimal
	Fld合計金額   decimal.Decimal
	Fld按分先    string
	Fld按分基準値  decimal.Decimal
	Fld按分誤差   decimal.Decimal // 按分結果に含まれている誤差
	Fld按分結果   decimal.Decimal
}
