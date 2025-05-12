package domain

import "github.com/shopspring/decimal"

type List按分結果 struct {
	List *OrderedMap[Key按分結果, decimal.Decimal]
}

func NewList按分結果() List按分結果 {
	return List按分結果{
		List: NewOrderedMap[Key按分結果, decimal.Decimal](),
	}
}

func (l *List按分結果) Add(key Key按分結果, 金額 decimal.Decimal) {
	if 金額合計, ok := l.List.Get(key); !ok {
		l.List.Set(key, 金額)
	} else {
		l.List.Set(key, 金額合計.Add(金額))
	}
}

func (l *List按分結果) Get() []*Ent按分結果 {
	按分結果一覧 := make([]*Ent按分結果, 0, l.List.Count())
	for _, k := range l.List.Keys() {
		合計金額, _ := l.List.Get(k)
		按分結果一覧 = append(按分結果一覧, &Ent按分結果{
			Fld計上年月:  k.Fld計上年月,
			Fld原価要素:  k.Fld原価要素,
			FldIs直接費: k.FldIs直接費,
			Fld借方税区分: k.Fld借方税区分,
			Fld借方税率:  k.Fld借方税率,
			Fld按分先:   k.Fld按分先,
			Fld金額:    合計金額,
		})
	}
	return 按分結果一覧
}
