package domain

import "github.com/shopspring/decimal"

type Key集計仕訳 struct {
	Fld計上年月   string
	Fld勘定科目   string
	Fldコストプール string
	Fld按分ルール1 string
	Fld按分ルール2 string
	Fld借方税区分  string
	Fld借方税率   decimal.Decimal
}
