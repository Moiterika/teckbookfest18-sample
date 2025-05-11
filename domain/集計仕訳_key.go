package domain

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Key集計仕訳 struct {
	Fld計上年月   string
	Fld原価要素   string
	Fldコストプール string
	Fld按分ルール1 string
	Fld按分ルール2 string
	Fld借方税区分  string
	Fld借方税率   decimal.Decimal
}

func newKey集計仕訳(e *Ent仕訳) (*Key集計仕訳, error) {
	if e.Val仕訳詳細 == nil {
		return nil, fmt.Errorf("仕訳詳細が未定義です。")
	}
	if e.Val仕訳詳細.Fld按分ルール1 != "対象外" {
		if e.Val仕訳詳細.Fld原価要素 == "" {
			return nil, fmt.Errorf("原価要素が未定義です。")
		}
		if e.Val仕訳詳細.Fldコストプール == "" {
			return nil, fmt.Errorf("コストプールが未定義です。")
		}
		if e.Val仕訳詳細.Fld按分ルール1 == "" {
			return nil, fmt.Errorf("按分ルール1が未定義です。")
		}
	}
	return &Key集計仕訳{
		Fld計上年月:   e.Val仕訳詳細.Fld計上年月,
		Fld原価要素:   e.Fld原価要素,
		Fldコストプール: e.Val仕訳詳細.Fldコストプール,
		Fld按分ルール1: e.Val仕訳詳細.Fld按分ルール1,
		Fld按分ルール2: e.Val仕訳詳細.Fld按分ルール2,
		Fld借方税区分:  e.Fld借方税区分,
		Fld借方税率:   e.Fld借方税率,
	}, nil
}
