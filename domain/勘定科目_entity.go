package domain

import "fmt"

type Enum基本ルール string

const (
	基本ルール_労務費配賦 Enum基本ルール = "労務費配賦"
	基本ルール_経費配賦  Enum基本ルール = "経費配賦"
	基本ルール_直課    Enum基本ルール = "直課"
	基本ルール_対象外   Enum基本ルール = "対象外"
)

type Ent勘定科目 struct {
	Fld勘定科目   string
	Fld基本ルール  Enum基本ルール
	Fldコストプール string // コストプール（任意項目）
}

func New基本ルール(基本ルール string) Enum基本ルール {
	switch 基本ルール {
	case string(基本ルール_労務費配賦):
		return 基本ルール_労務費配賦
	case string(基本ルール_経費配賦):
		return 基本ルール_経費配賦
	case string(基本ルール_直課):
		return 基本ルール_直課
	case string(基本ルール_対象外):
		return 基本ルール_対象外
	default:
		fmt.Printf("Invalid Enum基本ルール: %s\n", 基本ルール)
		return 基本ルール_対象外
	}
}
