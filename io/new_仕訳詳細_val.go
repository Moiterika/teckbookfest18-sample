package io

import (
	"database/sql"
	"teckbookfest18-sample/domain"
)

func newVal仕訳詳細FromDb(計上年月, コストプール, 按分ルール1, 按分ルール2 sql.NullString) *domain.Val仕訳詳細 {
	if !計上年月.Valid {
		return nil
	}
	return &domain.Val仕訳詳細{
		Fld計上年月:   計上年月.String,
		Fldコストプール: コストプール.String,
		Fld按分ルール1: 按分ルール1.String,
		Fld按分ルール2: 按分ルール2.String,
	}
}

func newVal仕訳詳細(計上年月, コストプール, 按分ルール1, 按分ルール2 string) *domain.Val仕訳詳細 {
	if 計上年月 != "" {
		return nil
	}
	return &domain.Val仕訳詳細{
		Fld計上年月:   計上年月,
		Fldコストプール: コストプール,
		Fld按分ルール1: 按分ルール1,
		Fld按分ルール2: 按分ルール2,
	}
}
