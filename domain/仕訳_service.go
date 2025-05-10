package domain

import (
	"fmt"
	"time"

	"github.com/Moiterika/a"
)

type Service仕訳 struct {
	csv  I仕訳CsvReader
	xlsx I仕訳XlsxIo
	repo I仕訳Repo
}

// NewService仕訳 は Service仕訳 のインスタンスを作成します
func NewService仕訳(csv I仕訳CsvReader, xlsx I仕訳XlsxIo, repo I仕訳Repo) *Service仕訳 {
	return &Service仕訳{
		csv:  csv,
		xlsx: xlsx,
		repo: repo,
	}
}

func (s *Service仕訳) Execute() error {
	// 1. CSVから仕訳データを読み取る
	csvRows, err := s.csv.ReadAll()
	if err != nil {
		return err
	}

	// 2. xlsxから仕訳データを読み取る
	xlsxRows, err := s.xlsx.Read仕訳一覧()
	if err != nil {
		return err
	}
	xlsxDic, err := a.ToMapWithErr(xlsxRows, func(e *Ent仕訳) 仕訳Key {
		return e.Key()
	})
	if err != nil {
		return err
	}

	// 3. xlsxの勘定科目を読み込む
	勘定科目一覧, err := s.xlsx.Read勘定科目一覧()
	if err != nil {
		return err
	}
	科目Dic, err := a.ToMapWithErr(勘定科目一覧, func(e *Ent勘定科目) string {
		return e.Fld勘定科目
	})
	if err != nil {
		return err
	}

	// 2. CSVで読み取った仕訳にxlsxの仕訳詳細をマージする
	for i, csvRow := range csvRows {
		// 計上年月
		t, err := time.Parse("2006/01/02", csvRow.Fld取引日) // YYYY/MM/DD
		if err != nil {
			return fmt.Errorf("CSV%d行目の取引日がYYYY/MM/DD形式ではありません。: %w", i+1, err)
		}
		計上年月 := t.Format("200601") // YYYYMM

		if x, ok := xlsxDic[csvRow.Key()]; ok {
			csvRow.GetVal仕訳詳細From(x) // xlsxにある仕訳詳細を取得してマージ
			// 取引日変更で計上年月が違っている場合があるので、その場合、警告
			fmt.Printf("【警告】仕訳一覧%d行目: 計上年月が取引日と違います。計上年月=%s、取引日=%s\n", i+1, 計上年月, csvRow.Fld取引日)
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
		}
	}

	// 3. 仕訳データをデータベースに保存する
	if err := s.repo.Save(csvRows); err != nil {
		return err
	}

	return nil

}
