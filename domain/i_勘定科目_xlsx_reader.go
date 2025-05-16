package domain

// I勘定科目XlsxReader は勘定科目データをxlsxから読み取るインターフェースです
type I勘定科目XlsxReader interface {
	// Read勘定科目一覧 は勘定科目データを読み取ります
	Read勘定科目一覧() ([]*Ent勘定科目, error)
}
