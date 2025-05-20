package io

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func getInt64Cell(row []string, i int) (int64, error) {
	if i > len(row)-1 {
		colName, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			return 0, fmt.Errorf("%d列目のxlsx列名変換に失敗しました。: %w", i+1, err)
		}
		return 0, fmt.Errorf("値がありません。列=%s", colName)
	}
	return strconv.ParseInt(row[i], 10, 64)
}
