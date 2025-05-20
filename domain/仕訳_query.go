package domain

// Qry仕訳 は仕訳データCSVを読み取るインターフェースです
type Qry仕訳 interface {
	// ReadAll は仕訳データを読み取ります
	ReadAll() ([]*Ent仕訳, error)
}
