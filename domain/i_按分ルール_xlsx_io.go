package domain

// I按分ルールXlsxIo は按分ルールデータxlsxを読み書きするインターフェースです
type I按分ルールXlsxIo interface {
	// Read按分ルール一覧 は按分ルールデータを読み取ります
	Read按分ルール一覧() ([]*Ent按分ルール, error)
	// Save は按分ルールデータを保存します
	Save([]*Ent按分ルール) error
}
