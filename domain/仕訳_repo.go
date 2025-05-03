package domain

// Repo仕訳 は仕訳データをデータベースに永続化するためのインターフェースです
type Repo仕訳 interface {
	// Save は仕訳データをデータベースに保存します
	Save(仕訳 []*Ent仕訳) error

	// FindAll はデータベースから全ての仕訳データを取得します
	FindAll() ([]*Ent仕訳, error)
}
