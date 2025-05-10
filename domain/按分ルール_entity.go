package domain

import "github.com/shopspring/decimal"

type Ent按分ルール struct {
	Fld按分ルール1 string
	Fld按分ルール2 string
	Fld按分先    string
	Fld按分基準値  decimal.Decimal
}

func (e *Ent按分ルール) Key() Key按分ルール {
	return Key按分ルール{
		Fld按分ルール1: e.Fld按分ルール1,
		Fld按分ルール2: e.Fld按分ルール2,
	}
}
