package domain

// I仕訳XlsxIo は仕訳データxlsxを読み書きするインターフェースです
type I仕訳XlsxIo interface {
	// Read仕訳一覧 は仕訳データを読み取ります
	Read仕訳一覧() ([]*Ent仕訳, error)
	// Save は仕訳データを保存します
	Save([]*Ent仕訳) error
}
