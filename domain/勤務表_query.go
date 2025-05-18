package domain

type Qry勤務表 interface {
	// Read勤務表 は勤務表データを読み取ります
	Read勤務表() ([]*Ent勤務表, error)
}
