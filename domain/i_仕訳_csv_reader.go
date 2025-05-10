package domain

// I仕訳CsvReader は仕訳データCSVを読み取るインターフェースです
type I仕訳CsvReader interface {
	// ReadAll は仕訳データを読み取ります
	ReadAll() ([]*Ent仕訳, error)
}
