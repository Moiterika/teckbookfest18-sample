package io

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
)

func getDecimalCell(row []string, i int) (decimal.Decimal, error) {
	if i > len(row)-1 {
		colName, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			return decimal.Zero, fmt.Errorf("%d列目のxlsx列名変換に失敗しました。: %w", i+1, err)
		}
		return decimal.Zero, fmt.Errorf("値がありません。列=%s", colName)
	}
	return convertToDecimal(row[i])
}
