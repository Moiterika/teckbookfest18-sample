package domain

// I仕訳Repo は仕訳データをデータベースに永続化するためのインターフェースです
type I仕訳Repo interface {
	// Save は仕訳データをデータベースに保存します
	Save(仕訳 []*Ent仕訳) error

	// FindAll はデータベースから全ての仕訳データを取得します
	FindAll() ([]*Ent仕訳, error)
}
