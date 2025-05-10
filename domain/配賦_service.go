package domain

import (
	"fmt"

	"github.com/Moiterika/a"
	"github.com/shopspring/decimal"
)

type Service配賦 struct {
	按分ルールxlsx  I按分ルールXlsxIo
	按分結果明細xlsx I按分結果明細XlsxWriter
}

// NewService配賦 は配賦サービスを生成します
func NewService配賦(xlsx I按分ルールXlsxIo, 按分結果明細xlsx I按分結果明細XlsxWriter) *Service配賦 {
	return &Service配賦{
		按分ルールxlsx:  xlsx,
		按分結果明細xlsx: 按分結果明細xlsx,
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

func (s *Service配賦) Execute配賦(集計仕訳一覧 []*Ent集計仕訳, 按分ルール一覧 []*Ent按分ルール) ([]*Ent按分結果明細, error) {
	// 1. 按分ルールをキーでマップ化
	按分ルールgroup, err := a.GroupByWithErr(按分ルール一覧, func(e *Ent按分ルール) Key按分ルール {
		return e.Key()
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	// 2. 集計仕訳データを按分ルールに基づき配賦する
	ret := make([]*Ent按分結果明細, 0)
	for _, e := range 集計仕訳一覧 {
		// 2-1. 直課の場合は按分しない
		if e.Fld按分ルール1 == "直課" {
			ret = append(ret, &Ent按分結果明細{
				// 集計仕訳の情報をセット
				Fld計上年月:   e.Fld計上年月,
				Fld勘定科目:   e.Fld勘定科目,
				Fldコストプール: e.Fldコストプール,
				Fld按分ルール1: e.Fld按分ルール1,
				Fld按分ルール2: e.Fld按分ルール2,
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
			return nil, fmt.Errorf("按分ルール一覧にない按分ルールです。%v", e)
		} else {
			// 2-3. 按分ルールに基づいて集計仕訳を按分する
			result, err := Calc按分(e.Fld合計金額, 按分基準, func(x *Ent按分ルール) decimal.Decimal {
				return x.Fld按分基準値
			})
			if err != nil {
				return nil, fmt.Errorf("按分計算でエラー: %w", err)
			}

			// 按分ルールに基づいて集計仕訳を按分する
			for _, r := range result {
				ret = append(ret, &Ent按分結果明細{
					// 集計仕訳の情報をセット
					Fld計上年月:   e.Fld計上年月,
					Fld勘定科目:   e.Fld勘定科目,
					Fldコストプール: e.Fldコストプール,
					Fld按分ルール1: e.Fld按分ルール1,
					Fld按分ルール2: e.Fld按分ルール2,
					Fld借方税区分:  e.Fld借方税区分,
					Fld借方税率:   e.Fld借方税率,
					Fld合計金額:   e.Fld合計金額,
					// 按分結果の情報をセット
					Fld按分先:   r.Original.Fld按分先,   // 按分ルールの按分先
					Fld按分基準値: r.Original.Fld按分基準値, // 按分ルールの按分基準値
					Fld按分誤差:  r.DiffValue,         // 按分誤差
					Fld按分結果:  r.AllocatedValue,    // 按分後の金額
				})
			}
		}
	}
	// 3. 按分結果をxlsxに保存する
	err = s.按分結果明細xlsx.Save(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
