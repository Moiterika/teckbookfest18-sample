package io

import (
	"strings"

	"github.com/shopspring/decimal"
)

// convertToDecimal は文字列をdecimal.Decimalに変換します
// 空文字の場合は0を返します
// エラーが発生した場合はエラーを返します
// 例: "123.45" -> decimal.NewFromFloat(123.45)
// 例: "" -> decimal.NewFromInt(0)
// 例: "123abc" -> error
func convertToDecimal(s string) (decimal.Decimal, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		// 空文字なら 0 を設定
		return decimal.NewFromInt(0), nil
	}
	return decimal.NewFromString(s)
}
