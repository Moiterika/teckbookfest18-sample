package domain

// I仕訳XlsxIo は仕訳データxlsxを読み書きするインターフェースです
type I仕訳XlsxIo interface {
	// Read勘定科目一覧 は勘定科目データを読み取ります
	Read勘定科目一覧() ([]*Ent勘定科目, error)
	// Read仕訳一覧 は仕訳データを読み取ります
	Read仕訳一覧() ([]*Ent仕訳, error)
	// Save は仕訳データを保存します
	Save([]*Ent仕訳) error
}
