package domain

import "github.com/shopspring/decimal"

type Ent勤務表 struct {
	Fld作業内容       string          // コストプール（原価部門 or 指図）
	Fld作業時間_分     decimal.Decimal // 作業時間（分）
	Fld労務費按分用の計上月 string
	Fld経費按分用の計上月  string
}
