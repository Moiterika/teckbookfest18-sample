package io

func getStringCell(row []string, i int) string {
	if i > len(row)-1 {
		return ""
	}
	return row[i]
}
