package domain

type Key按分ルール struct {
	Fld按分ルール1 string
	Fld按分ルール2 string
}

func newKey按分ルール(e *Ent集計仕訳) Key按分ルール {
	return Key按分ルール{
		Fld按分ルール1: e.Fld按分ルール1,
		Fld按分ルール2: e.Fld按分ルール2,
	}
}
