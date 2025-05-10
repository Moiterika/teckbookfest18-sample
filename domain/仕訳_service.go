package domain

import (
	"fmt"
	"time"

	"github.com/Moiterika/a"
)

type Service仕訳 struct {
	csv  I仕訳CsvReader
	xlsx I仕訳XlsxIo
}

// NewService仕訳 は Service仕訳 のインスタンスを作成します
func NewService仕訳(csv I仕訳CsvReader, xlsx I仕訳XlsxIo) *Service仕訳 {
	return &Service仕訳{
		csv:  csv,
		xlsx: xlsx,
	}
}

var Error未定義仕訳 = fmt.Errorf("按分ルールが未定義です")

func (s *Service仕訳) Query() ([]*Ent仕訳, error) {
	// 1. CSVから仕訳データを読み取る
	csvRows, err := s.csv.ReadAll()
	if err != nil {
		err = fmt.Errorf("csvの読込でエラー: %w", err)
		fmt.Printf("%v\n", err)
		return nil, err
	}

	// 2. xlsxから仕訳データを読み取る
	xlsxRows, err := s.xlsx.Read仕訳一覧()
	if err != nil {
		err = fmt.Errorf("xlsxの「仕訳一覧」読込でエラー: %w", err)
		fmt.Printf("%v\n", err)
		return nil, err
	}
	xlsxDic, err := a.ToMapWithErr(xlsxRows, func(e *Ent仕訳) key仕訳 {
		return e.Key()
	})
	if err != nil {
		err = fmt.Errorf("xlsxの「仕訳一覧」のキーが重複: %w", err)
		fmt.Printf("%v\n", err)
		return nil, err
	}

	// 3. xlsxの勘定科目を読み込む
	勘定科目一覧, err := s.xlsx.Read勘定科目一覧()
	if err != nil {
		err = fmt.Errorf("xlsxの「勘定科目一覧」読込でエラー: %w", err)
		fmt.Printf("%v\n", err)
		return nil, err
	}
	科目Dic, err := a.ToMapWithErr(勘定科目一覧, func(e *Ent勘定科目) string {
		return e.Fld勘定科目
	})
	if err != nil {
		err = fmt.Errorf("xlsxの「勘定科目一覧」のキーが重複: %w", err)
		fmt.Printf("%v\n", err)
		return nil, err
	}

	// 2. CSVで読み取った仕訳にxlsxの仕訳詳細をマージする
	var hasErr bool
	for i, csvRow := range csvRows {
		// 計上年月
		t, err := time.Parse("2006/01/02", csvRow.Fld取引日) // YYYY/MM/DD
		if err != nil {
			err = fmt.Errorf("CSV%d行目の取引日がYYYY/MM/DD形式ではありません。: %w", i+1, err)
			fmt.Printf("%v\n", err)
			return nil, err
		}
		計上年月 := t.Format("200601") // YYYYMM

		if x, ok := xlsxDic[csvRow.Key()]; ok {
			csvRow.GetVal仕訳詳細From(x) // xlsxにある仕訳詳細を取得してマージ
			// 取引日変更で計上年月が違っている場合があるので、その場合、警告
			fmt.Printf("【警告】仕訳一覧%d行目:計上年月が取引日と違います。計上年月=%s、取引日=%s\n", i+1, 計上年月, csvRow.Fld取引日)
			continue
		}

		if 科目, ok := 科目Dic[csvRow.Fld借方勘定科目]; ok {
			コストプール := 科目.Fldコストプール
			if 科目.Fld基本ルール == 基本ルール_直課 && csvRow.Fld借方部門 != "" {
				コストプール = csvRow.Fld借方部門
			}
			按分ルール2 := ""
			switch 科目.Fld基本ルール {
			case 基本ルール_労務費配賦:
				fallthrough
			case 基本ルール_経費配賦:
				按分ルール2 = 計上年月
			}

			csvRow.Val仕訳詳細 = &Val仕訳詳細{
				Fld計上年月:   計上年月,
				Fldコストプール: コストプール,
				Fld按分ルール1: string(科目.Fld基本ルール),
				Fld按分ルール2: 按分ルール2,
			}
			continue
		}
		hasErr = true
		fmt.Printf("仕訳一覧%d行目: 按分ルールが未定義です。\n", i+1)
	}
	if hasErr {
		return csvRows, Error未定義仕訳
	}
	return csvRows, nil
}
