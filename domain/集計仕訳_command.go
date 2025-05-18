package domain

// Cmd集計仕訳 は集計仕訳データxlsxを書き出すインターフェースです
// Save は集計仕訳一覧をxlsxファイルに書き出します
// 戻り値のerrorは書き出し時のエラーを表します
type Cmd集計仕訳 interface {
	Save([]*Ent集計仕訳) error
}
