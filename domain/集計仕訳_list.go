package domain

import "github.com/shopspring/decimal"

type List集計仕訳 struct {
	List *OrderedMap[Key集計仕訳, decimal.Decimal]
}

func NewList集計仕訳() List集計仕訳 {
	return List集計仕訳{
		List: NewOrderedMap[Key集計仕訳, decimal.Decimal](),
	}
}

func (l *List集計仕訳) Add(key Key集計仕訳, 金額 decimal.Decimal) {
	if 金額合計, ok := l.List.Get(key); !ok {
		l.List.Set(key, 金額)
	} else {
		l.List.Set(key, 金額合計.Add(金額))
	}
}

func (l *List集計仕訳) Get() []*Ent集計仕訳 {
	集計仕訳一覧 := make([]*Ent集計仕訳, 0, l.List.Count())
	for _, key := range l.List.Keys() {
		合計金額, _ := l.List.Get(key)
		集計仕訳一覧 = append(集計仕訳一覧, &Ent集計仕訳{
			Fld計上年月:   key.Fld計上年月,
			Fld原価要素:   key.Fld原価要素,
			Fldコストプール: key.Fldコストプール,
			Fld按分ルール1: key.Fld按分ルール1,
			Fld按分ルール2: key.Fld按分ルール2,
			Fld借方税区分:  key.Fld借方税区分,
			Fld借方税率:   key.Fld借方税率,
			Fld合計金額:   合計金額,
		})
	}
	return 集計仕訳一覧
}
