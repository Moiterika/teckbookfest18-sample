package domain

type key工数集計 struct {
	Fld按分ルール1 string
	Fld按分ルール2 string
	Fld按分先    string
}

func newKey工数集計For労務費(e *Ent勤務表) key工数集計 {
	return key工数集計{
		Fld按分ルール1: "労務費配賦",
		Fld按分ルール2: e.Fld労務費按分用の計上月,
		Fld按分先:    e.Fld作業内容,
	}
}

func newKey工数集計For経費(e *Ent勤務表) key工数集計 {
	return key工数集計{
		Fld按分ルール1: "経費配賦",
		Fld按分ルール2: e.Fld経費按分用の計上月,
		Fld按分先:    e.Fld作業内容,
	}
}
