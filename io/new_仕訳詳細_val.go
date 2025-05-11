package io

import (
	"teckbookfest18-sample/domain"
)

func newVal仕訳詳細(計上年月, 原価要素, コストプール, 按分ルール1, 按分ルール2 string) *domain.Val仕訳詳細 {
	if 計上年月 == "" {
		return nil
	}
	return &domain.Val仕訳詳細{
		Fld計上年月:   計上年月,
		Fld原価要素:   原価要素,
		Fldコストプール: コストプール,
		Fld按分ルール1: 按分ルール1,
		Fld按分ルール2: 按分ルール2,
	}
}
