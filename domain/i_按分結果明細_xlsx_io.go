package domain

// Ent按分結果明細をxlsxで読み書きするインターフェース
type I按分結果明細XlsxIo interface {
	// // 按分結果明細一覧を全件読み取る
	// Read() ([]*Ent按分結果明細, error)
	// Save は按分結果明細一覧を保存する
	Save([]*Ent按分結果明細) error
}
