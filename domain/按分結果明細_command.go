package domain

// Ent按分結果明細をxlsxで保存するためのインターフェース
// 按分結果明細の保存処理のみを定義する
// ファイルの読み込み処理は含まない
// 実装はio層で行うこと
type Cmd按分結果明細 interface {
	// Saveは按分結果明細一覧をxlsxファイルに保存する
	Save([]*Ent按分結果明細) error
}
