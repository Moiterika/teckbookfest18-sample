package domain

// I集計仕訳XlsxWriter は集計仕訳データxlsxを書き出すインターフェースです
// Save は集計仕訳一覧をxlsxファイルに書き出します
// 戻り値のerrorは書き出し時のエラーを表します
type I集計仕訳XlsxWriter interface {
	Save([]*Ent集計仕訳) error
}
