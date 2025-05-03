package domain

// Query仕訳 は仕訳データCSVを読み取るインターフェースです
type Query仕訳 interface {
	// Read は仕訳データを読み取ります
	Read() ([]*Ent仕訳, error)
}
