package domain

import (
	"sort"

	"github.com/shopspring/decimal"
)

type List按分結果 struct {
	Map      map[string]decimal.Decimal // 文字列化したキーと金額のマップ
	KeyStore map[string]*Key按分結果        // キーのハッシュとキー構造体のマッピング
}

func NewList按分結果() List按分結果 {
	return List按分結果{
		Map:      make(map[string]decimal.Decimal),
		KeyStore: make(map[string]*Key按分結果),
	}
}

func (l *List按分結果) Add(key Key按分結果, 金額 decimal.Decimal) {
	// キーを文字列にハッシュ化
	hashKey := key.Hash()

	// キーの実体を保存
	l.KeyStore[hashKey] = &key

	// 金額を集計
	if 金額合計, ok := l.Map[hashKey]; ok {
		l.Map[hashKey] = 金額合計.Add(金額)
	} else {
		l.Map[hashKey] = 金額
	}
}

func (l *List按分結果) Get() []*Ent按分結果 {
	// ソート用にキーのスライスを作成
	keys := make([]string, 0, len(l.Map))
	for k := range l.Map {
		keys = append(keys, k)
	}

	// キーをソート
	sort.Strings(keys)

	// 結果スライスを作成
	按分結果一覧 := make([]*Ent按分結果, 0, len(l.Map))

	// ソートされたキーの順に按分結果を作成
	for _, hashKey := range keys {
		k := l.KeyStore[hashKey]
		合計金額 := l.Map[hashKey]

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
