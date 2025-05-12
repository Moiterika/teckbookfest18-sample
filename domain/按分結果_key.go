package domain

import "github.com/shopspring/decimal"

type Key按分結果 struct {
	Fld計上年月  string
	Fld原価要素  string
	FldIs直接費 bool
	Fld借方税区分 string
	Fld借方税率  decimal.Decimal
	Fld按分先   string
}

func newKey按分結果(e *Ent按分結果明細) Key按分結果 {
	return Key按分結果{
		Fld計上年月:  e.Fld計上年月,
		Fld原価要素:  e.Fld原価要素,
		FldIs直接費: e.FldIs直接費,
		Fld借方税区分: e.Fld借方税区分,
		Fld借方税率:  e.Fld借方税率,
		Fld按分先:   e.Fld按分先,
	}
}
