package domain

import (
	"sort"

	"github.com/shopspring/decimal"
)

type List集計仕訳 struct {
	Map      map[string]decimal.Decimal // 文字列化したキーと金額のマップ
	KeyStore map[string]*Key集計仕訳        // キーのハッシュとキー構造体のマッピング
}

func NewList集計仕訳() List集計仕訳 {
	return List集計仕訳{
		Map:      make(map[string]decimal.Decimal),
		KeyStore: make(map[string]*Key集計仕訳),
	}
}

func (l *List集計仕訳) Add(key Key集計仕訳, 金額 decimal.Decimal) {
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

func (l *List集計仕訳) Get() []*Ent集計仕訳 {
	// ソート用にキーのスライスを作成
	keys := make([]string, 0, len(l.Map))
	for k := range l.Map {
		keys = append(keys, k)
	}

	// キーをソート
	sort.Strings(keys)

	// 結果スライスを作成
	集計仕訳一覧 := make([]*Ent集計仕訳, 0, len(l.Map))

	// ソートされたキーの順に集計仕訳を作成
	for _, hashKey := range keys {
		key := l.KeyStore[hashKey]
		金額 := l.Map[hashKey]

		集計仕訳一覧 = append(集計仕訳一覧, &Ent集計仕訳{
			Fld計上年月:   key.Fld計上年月,
			Fld原価要素:   key.Fld原価要素,
			Fldコストプール: key.Fldコストプール,
			Fld按分ルール1: key.Fld按分ルール1,
			Fld按分ルール2: key.Fld按分ルール2,
			Fld借方税区分:  key.Fld借方税区分,
			Fld借方税率:   key.Fld借方税率,
			Fld合計金額:   金額,
		})
	}

	return 集計仕訳一覧
}
