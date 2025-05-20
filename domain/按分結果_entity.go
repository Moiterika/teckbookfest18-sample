package domain

import (
	"strings"

	"github.com/shopspring/decimal"
)

type Ent按分結果 struct {
	Fld計上年月  string
	Fld按分先   string
	Fld原価要素  string
	FldIs直接費 bool
	Fld借方税区分 string
	Fld借方税率  decimal.Decimal
	Fld金額    decimal.Decimal
}

func (e *Ent按分結果) Calc税() decimal.Decimal {
	if e.Fld借方税率.IsZero() {
		return decimal.Zero
	}
	h := decimal.NewFromInt32(100)
	if strings.Contains(e.Fld借方税区分, "控80") {
		k := decimal.NewFromInt32(2)
		return e.Fld金額.Mul(e.Fld借方税率.Add(k.Neg())).Div(h.Add(k)).RoundDown(0)
	}
	return e.Fld金額.Mul(e.Fld借方税率).Div(h).RoundDown(0)
}

func (e *Ent按分結果) Calc税込金額() decimal.Decimal {
	return e.Fld金額.Add(e.Calc税())
}
