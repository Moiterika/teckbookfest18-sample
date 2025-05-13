package domain

import "github.com/shopspring/decimal"

type Ent集計仕訳 struct {
	Fld計上年月   string
	Fld原価要素   string
	Fldコストプール string
	Fld按分ルール1 string
	Fld按分ルール2 string
	Fld借方税区分  string
	Fld借方税率   decimal.Decimal
	Fld合計金額   decimal.Decimal
}
