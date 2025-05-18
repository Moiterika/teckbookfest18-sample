package domain

import (
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// Mock工数Qry勤務表 は Qry勤務表 のモック（工数集計用）
type Mock工数Qry勤務表 struct {
	勤務表一覧 []*Ent勤務表
	Err   error
}

func (m *Mock工数Qry勤務表) Read勤務表() ([]*Ent勤務表, error) {
	return m.勤務表一覧, m.Err
}

// Mock工数Rep按分ルール は Rep按分ルール のモック（工数集計用）
type Mock工数Rep按分ルール struct {
	按分ルール一覧   []*Ent按分ルール
	SavedData []*Ent按分ルール
	Err       error
}

func (m *Mock工数Rep按分ルール) Read按分ルール一覧() ([]*Ent按分ルール, error) {
	return m.按分ルール一覧, nil
}

func (m *Mock工数Rep按分ルール) Save(data []*Ent按分ルール) error {
	if m.Err != nil {
		return m.Err
	}
	m.SavedData = data
	return nil
}

func TestNewService工数集計(t *testing.T) {
	// テスト用のモックを準備
	mockQry勤務表 := &Mock工数Qry勤務表{}
	mockRep按分ルール := &Mock工数Rep按分ルール{}

	// テスト対象のサービスを作成
	service := NewService工数集計(mockQry勤務表, mockRep按分ルール)

	// 期待通りの依存性が注入されていることを検証
	assert.Equal(t, mockQry勤務表, service.勤務表Reader)
	assert.Equal(t, mockRep按分ルール, service.按分ルールIo)
}

func TestExecute工数集計_正常系(t *testing.T) {
	// テストデータの準備
	勤務表一覧 := []*Ent勤務表{
		{
			Fld作業内容:       "プロジェクトA",
			Fld作業時間_分:     decimal.NewFromInt(120),
			Fld労務費按分用の計上月: "202301",
			Fld経費按分用の計上月:  "202301",
		},
		{
			Fld作業内容:       "プロジェクトA",
			Fld作業時間_分:     decimal.NewFromInt(60),
			Fld労務費按分用の計上月: "202301",
			Fld経費按分用の計上月:  "202301",
		},
		{
			Fld作業内容:       "プロジェクトB",
			Fld作業時間_分:     decimal.NewFromInt(180),
			Fld労務費按分用の計上月: "202302",
			Fld経費按分用の計上月:  "202302",
		},
	}

	mockQry勤務表 := &Mock工数Qry勤務表{
		勤務表一覧: 勤務表一覧,
	}
	mockRep按分ルール := &Mock工数Rep按分ルール{}

	service := NewService工数集計(mockQry勤務表, mockRep按分ルール)

	// テスト対象メソッドを実行
	err := service.Execute工数集計()

	// 検証
	assert.NoError(t, err)
	assert.NotNil(t, mockRep按分ルール.SavedData)

	// 実際に作成される按分ルールのチェック
	assert.Equal(t, 4, len(mockRep按分ルール.SavedData))

	// プロジェクトAの労務費工数が正しく集計されていることを確認（120 + 60 = 180分）
	// データがソートされるので、期待する順序で確認
	assert.Equal(t, "労務費配賦", mockRep按分ルール.SavedData[0].Fld按分ルール1)
	assert.Equal(t, "202301", mockRep按分ルール.SavedData[0].Fld按分ルール2)
	assert.Equal(t, "プロジェクトA", mockRep按分ルール.SavedData[0].Fld按分先)
	assert.Equal(t, decimal.NewFromInt(180), mockRep按分ルール.SavedData[0].Fld按分基準値)

	// プロジェクトBの労務費工数が正しく集計されていることを確認
	assert.Equal(t, "労務費配賦", mockRep按分ルール.SavedData[1].Fld按分ルール1)
	assert.Equal(t, "202302", mockRep按分ルール.SavedData[1].Fld按分ルール2)
	assert.Equal(t, "プロジェクトB", mockRep按分ルール.SavedData[1].Fld按分先)
	assert.Equal(t, decimal.NewFromInt(180), mockRep按分ルール.SavedData[1].Fld按分基準値)

	// 経費工数も確認（ソートされるので労務費の後）
	assert.Equal(t, "経費配賦", mockRep按分ルール.SavedData[2].Fld按分ルール1)
	assert.Equal(t, "202301", mockRep按分ルール.SavedData[2].Fld按分ルール2)
	assert.Equal(t, "プロジェクトA", mockRep按分ルール.SavedData[2].Fld按分先)
	assert.Equal(t, decimal.NewFromInt(180), mockRep按分ルール.SavedData[2].Fld按分基準値)

	// 別のプロジェクトの経費工数も確認
	assert.Equal(t, "経費配賦", mockRep按分ルール.SavedData[3].Fld按分ルール1)
	assert.Equal(t, "202302", mockRep按分ルール.SavedData[3].Fld按分ルール2)
	assert.Equal(t, "プロジェクトB", mockRep按分ルール.SavedData[3].Fld按分先)
	assert.Equal(t, decimal.NewFromInt(180), mockRep按分ルール.SavedData[3].Fld按分基準値)
}

func TestExecute工数集計_勤務表読み込みエラー(t *testing.T) {
	// エラーケースのテスト
	mockQry勤務表 := &Mock工数Qry勤務表{
		Err: errors.New("勤務表の読み込みエラー"),
	}
	mockRep按分ルール := &Mock工数Rep按分ルール{}

	service := NewService工数集計(mockQry勤務表, mockRep按分ルール)

	// テスト対象メソッドを実行
	err := service.Execute工数集計()

	// エラーが発生することを確認
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "勤務表の読み込みに失敗")
	assert.Nil(t, mockRep按分ルール.SavedData) // 保存処理は実行されていないはず
}

func TestExecute工数集計_按分ルール保存エラー(t *testing.T) {
	// テストデータの準備
	勤務表一覧 := []*Ent勤務表{
		{
			Fld作業内容:       "プロジェクトA",
			Fld作業時間_分:     decimal.NewFromInt(120),
			Fld労務費按分用の計上月: "202301",
			Fld経費按分用の計上月:  "202301",
		},
	}

	mockQry勤務表 := &Mock工数Qry勤務表{
		勤務表一覧: 勤務表一覧,
	}
	mockRep按分ルール := &Mock工数Rep按分ルール{
		Err: errors.New("按分ルールの保存エラー"),
	}

	service := NewService工数集計(mockQry勤務表, mockRep按分ルール)

	// テスト対象メソッドを実行
	err := service.Execute工数集計()

	// エラーが発生することを確認
	assert.Error(t, err)
	assert.Equal(t, "按分ルールの保存エラー", err.Error())
}
