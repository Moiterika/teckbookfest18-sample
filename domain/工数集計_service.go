package domain

import (
	"fmt"
	"sort"

	"github.com/shopspring/decimal"
)

// Service工数集計 は工数集計関連のサービスを提供します
type Service工数集計 struct {
	勤務表Reader Qry勤務表
	按分ルールIo   Rep按分ルール
}

// NewService工数集計 は Service工数集計 のインスタンスを作成します
func NewService工数集計(勤務表Reader Qry勤務表, 按分ルールIo Rep按分ルール) *Service工数集計 {
	return &Service工数集計{
		勤務表Reader: 勤務表Reader,
		按分ルールIo:   按分ルールIo,
	}
}

// Execute工数集計 は工数集計処理を実行します
// 勤務表を読み込み、按分ルールを生成して保存する一連の処理を行います
func (s *Service工数集計) Execute工数集計() error {
	// 勤務表データを読み取る
	勤務表一覧, err := s.勤務表Reader.Read勤務表()
	if err != nil {
		err = fmt.Errorf("勤務表の読み込みに失敗: %w", err)
		return err
	}
	労務費工数 := make(map[key工数集計]decimal.Decimal)
	for _, e := range 勤務表一覧 {
		// 労務費工数を集計
		k := newKey工数集計For労務費(e)
		if v, ok := 労務費工数[k]; !ok {
			労務費工数[k] = e.Fld作業時間_分
		} else {
			労務費工数[k] = v.Add(e.Fld作業時間_分)
		}
	}
	経費工数 := make(map[key工数集計]decimal.Decimal)
	for _, e := range 勤務表一覧 {
		// 労務費工数を集計
		k := newKey工数集計For経費(e)
		if v, ok := 経費工数[k]; !ok {
			経費工数[k] = e.Fld作業時間_分
		} else {
			経費工数[k] = v.Add(e.Fld作業時間_分)
		}
	}

	// 按分ルール一覧に追加
	按分ルール一覧 := make([]*Ent按分ルール, 0)
	for k, v := range 労務費工数 {
		按分ルール一覧 = append(按分ルール一覧, &Ent按分ルール{
			Fld按分ルール1: k.Fld按分ルール1,
			Fld按分ルール2: k.Fld按分ルール2,
			Fld按分先:    k.Fld按分先,
			Fld按分基準値:  v,
		})

	}
	for k, v := range 経費工数 {
		按分ルール一覧 = append(按分ルール一覧, &Ent按分ルール{
			Fld按分ルール1: k.Fld按分ルール1,
			Fld按分ルール2: k.Fld按分ルール2,
			Fld按分先:    k.Fld按分先,
			Fld按分基準値:  v,
		})
	}

	sort.SliceStable(按分ルール一覧, func(i, j int) bool {
		// 按分ルール1, 按分ルール2, 按分先の順にソート
		switch {
		case 按分ルール一覧[i].Fld按分ルール1 != 按分ルール一覧[j].Fld按分ルール1:
			return 按分ルール一覧[i].Fld按分ルール1 < 按分ルール一覧[j].Fld按分ルール1
		case 按分ルール一覧[i].Fld按分ルール2 != 按分ルール一覧[j].Fld按分ルール2:
			return 按分ルール一覧[i].Fld按分ルール2 < 按分ルール一覧[j].Fld按分ルール2
		case 按分ルール一覧[i].Fld按分先 != 按分ルール一覧[j].Fld按分先:
			return 按分ルール一覧[i].Fld按分先 < 按分ルール一覧[j].Fld按分先
		default:
			return false
		}
	})

	// 按分ルールを保存
	return s.按分ルールIo.Save(按分ルール一覧)
}
