package domain

type I勤務表XlsxReader interface {
	// ReadAll は勤務表データを読み取ります
	ReadAll() ([]*Ent勤務表, error)
}
