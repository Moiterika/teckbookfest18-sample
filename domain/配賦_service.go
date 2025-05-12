package domain

import (
	"fmt"
	"sort"

	"github.com/Moiterika/a"
	"github.com/shopspring/decimal"
)

type Service配賦 struct {
	按分ルールxlsx  I按分ルールXlsxIo
	按分結果明細xlsx I按分結果明細XlsxWriter
	按分結果xlsx   I按分結果XlsxWriter
}

// NewService配賦 は配賦サービスを生成します
func NewService配賦(xlsx I按分ルールXlsxIo, 按分結果明細xlsx I按分結果明細XlsxWriter, 按分結果xlsx I按分結果XlsxWriter) *Service配賦 {
	return &Service配賦{
		按分ルールxlsx:  xlsx,
		按分結果明細xlsx: 按分結果明細xlsx,
		按分結果xlsx:   按分結果xlsx,
	}
}

// Query按分ルール一覧 は按分ルールデータを読み取ります
func (s *Service配賦) Query按分ルール一覧() ([]*Ent按分ルール, error) {
	// 1. xlsxから按分ルールデータを読み取る
	xlsxRows, err := s.按分ルールxlsx.Read按分ルール一覧()
	if err != nil {
		err = fmt.Errorf("xlsxの「按分ルール」読込でエラー: %w", err)
		fmt.Printf("%v\n", err)
		return nil, err
	}
	return xlsxRows, nil
}

func (s *Service配賦) Execute配賦(集計仕訳一覧 []*Ent集計仕訳, 按分ルール一覧 []*Ent按分ルール) error {
	// 1. 按分ルールをキーでマップ化
	按分ルールgroup, err := a.GroupByWithErr(按分ルール一覧, func(e *Ent按分ルール) Key按分ルール {
		return e.Key()
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	// 2. 集計仕訳データを按分ルールに基づき配賦する
	按分結果明細一覧 := make([]*Ent按分結果明細, 0)
	for _, e := range 集計仕訳一覧 {
		// 2-1. 直課の場合は按分しない
		if e.Fld按分ルール1 == "直課" {
			按分結果明細一覧 = append(按分結果明細一覧, &Ent按分結果明細{
				// 集計仕訳の情報をセット
				Fld計上年月:   e.Fld計上年月,
				Fld原価要素:   e.Fld原価要素,
				Fldコストプール: e.Fldコストプール,
				Fld按分ルール1: e.Fld按分ルール1,
				Fld按分ルール2: e.Fld按分ルール2,
				FldIs直接費:  true,
				Fld借方税区分:  e.Fld借方税区分,
				Fld借方税率:   e.Fld借方税率,
				Fld合計金額:   e.Fld合計金額,
				// 按分結果の情報をセット
				Fld按分先:   e.Fldコストプール,           // 按分先はコストプール
				Fld按分基準値: decimal.NewFromInt(1), // 按分基準値は1
				Fld按分誤差:  decimal.NewFromInt(0), // 按分誤差は0
				Fld按分結果:  e.Fld合計金額,             // 按分後の金額はそのまま
			})
			continue
		}

		// 2-2. 按分ルールが未定義の場合はエラー
		if 按分基準, ok := 按分ルールgroup[newKey按分ルール(e)]; !ok {
			// 按分ルールが未定義
			return fmt.Errorf("按分ルール一覧にない按分ルールです。%v", e)
		} else {
			// 2-3. 按分ルールに基づいて集計仕訳を按分する
			result, err := Calc按分(e.Fld合計金額, 按分基準, func(x *Ent按分ルール) decimal.Decimal {
				return x.Fld按分基準値
			})
			if err != nil {
				return fmt.Errorf("按分計算でエラー: %w", err)
			}

			// 按分ルールに基づいて集計仕訳を按分する
			for _, r := range result {
				按分結果明細一覧 = append(按分結果明細一覧, &Ent按分結果明細{
					// 集計仕訳の情報をセット
					Fld計上年月:   e.Fld計上年月,
					Fld原価要素:   e.Fld原価要素,
					Fldコストプール: e.Fldコストプール,
					Fld按分ルール1: e.Fld按分ルール1,
					Fld按分ルール2: e.Fld按分ルール2,
					FldIs直接費:  false,
					Fld借方税区分:  e.Fld借方税区分,
					Fld借方税率:   e.Fld借方税率,
					Fld合計金額:   e.Fld合計金額,
					// 按分結果の情報をセット
					Fld按分先:   r.Original.Fld按分先,   // 挹分ルールの按分先
					Fld按分基準値: r.Original.Fld按分基準値, // 挹分ルールの按分基準値
					Fld按分誤差:  r.DiffValue,         // 按分誤差
					Fld按分結果:  r.AllocatedValue,    // 挹分後の金額
				})
			}
		}
	}
	// 3. 按分結果明細の一覧をxlsxに保存する
	err = s.按分結果明細xlsx.Save(按分結果明細一覧)
	if err != nil {
		return err
	}

	// 4. 按分結果明細を集約する
	按分結果一覧 := NewList按分結果()
	for _, e := range 按分結果明細一覧 {
		key := newKey按分結果(e)
		按分結果一覧.Add(key, e.Fld按分結果)
	}
	ret := 按分結果一覧.Get()
	sort.SliceStable(ret, func(i, j int) bool {
		// 各フィールドをswitch文で順次比較
		switch {
		case ret[i].Fld計上年月 != ret[j].Fld計上年月:
			return ret[i].Fld計上年月 < ret[j].Fld計上年月
		case ret[i].Fld按分先 != ret[j].Fld按分先:
			return ret[i].Fld按分先 < ret[j].Fld按分先
		case ret[i].Fld原価要素 != ret[j].Fld原価要素:
			return ret[i].Fld原価要素 < ret[j].Fld原価要素
		case ret[i].FldIs直接費 != ret[j].FldIs直接費:
			return !ret[i].FldIs直接費 // i番目がfalse, j番目がtrueの順
		case !ret[i].Fld金額.Equal(ret[j].Fld金額):
			return ret[i].Fld金額.GreaterThan(ret[j].Fld金額) // 降順
		default:
			return false
		}
	})

	// 5. 按分結果の一覧をxlsxに保存する
	return s.按分結果xlsx.Save(ret)
}
