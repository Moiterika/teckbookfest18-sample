package domain

import (
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// モックの作成
type MockQry仕訳 struct {
	mock.Mock
}

func (m *MockQry仕訳) ReadAll() ([]*Ent仕訳, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Ent仕訳), args.Error(1)
}

type MockRep仕訳 struct {
	mock.Mock
}

func (m *MockRep仕訳) Read仕訳一覧() ([]*Ent仕訳, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Ent仕訳), args.Error(1)
}

func (m *MockRep仕訳) Save(仕訳一覧 []*Ent仕訳) error {
	args := m.Called(仕訳一覧)
	return args.Error(0)
}

type MockQry勘定科目 struct {
	mock.Mock
}

func (m *MockQry勘定科目) Read勘定科目一覧() ([]*Ent勘定科目, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Ent勘定科目), args.Error(1)
}

// テスト用のヘルパー関数
func createSample仕訳() *Ent仕訳 {
	return &Ent仕訳{
		FldNo:     1,
		Fld取引日:    "2024/05/01",
		Fld管理番号:   "123",
		Fld借方勘定科目: "旅費交通費",
		Fld借方金額:   decimal.NewFromInt(10000),
		Fld借方税区分:  "課税",
		Fld借方税金額:  decimal.NewFromInt(1000),
		Fld借方部門:   "営業部",
	}
}

func createSample勘定科目(勘定科目 string, 基本ルール Enum基本ルール) *Ent勘定科目 {
	return &Ent勘定科目{
		Fld勘定科目:   勘定科目,
		Fld原価要素:   "経費",
		Fld基本ルール:  基本ルール,
		Fldコストプール: "本社",
	}
}

// テスト
func TestNewService仕訳(t *testing.T) {
	// モックの作成
	mockCSV := new(MockQry仕訳)
	mock仕訳XLSX := new(MockRep仕訳)
	mock勘定科目XLSX := new(MockQry勘定科目)

	// テスト実行
	service := NewService仕訳(mockCSV, mock仕訳XLSX, mock勘定科目XLSX)

	// 検証
	assert.NotNil(t, service)
	assert.Same(t, mockCSV, service.csv)
	assert.Same(t, mock仕訳XLSX, service.仕訳xlsx)
	assert.Same(t, mock勘定科目XLSX, service.勘定科目xlsx)
}

func TestExecute仕訳集計_正常系(t *testing.T) {
	// モック準備
	mockCSV := new(MockQry仕訳)
	mock仕訳XLSX := new(MockRep仕訳)
	mock勘定科目XLSX := new(MockQry勘定科目)

	// CSVの仕訳データ
	csv仕訳 := []*Ent仕訳{
		createSample仕訳(),
	}

	// XLSXの仕訳データ（既存データ）
	xlsx仕訳 := []*Ent仕訳{
		func() *Ent仕訳 {
			j := createSample仕訳()
			j.Val仕訳詳細 = &Val仕訳詳細{
				Fld計上年月:   "202405",
				Fld原価要素:   "経費",
				Fldコストプール: "営業部",
				Fld按分ルール1: string(基本ルール_経費配賦),
				Fld按分ルール2: "202405",
			}
			return j
		}(),
	}

	// 勘定科目データ
	勘定科目一覧 := []*Ent勘定科目{
		createSample勘定科目("旅費交通費", 基本ルール_経費配賦),
	}

	// モックの動作設定
	mockCSV.On("ReadAll").Return(csv仕訳, nil)
	mock仕訳XLSX.On("Read仕訳一覧").Return(xlsx仕訳, nil)
	mock仕訳XLSX.On("Save", mock.Anything).Return(nil)
	mock勘定科目XLSX.On("Read勘定科目一覧").Return(勘定科目一覧, nil)

	// テスト対象
	service := NewService仕訳(mockCSV, mock仕訳XLSX, mock勘定科目XLSX)
	result, err := service.Execute仕訳集計()

	// 検証
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// モックの呼び出し検証
	mockCSV.AssertExpectations(t)
	mock仕訳XLSX.AssertExpectations(t)
	mock勘定科目XLSX.AssertExpectations(t)
}

func TestExecute仕訳集計_CSV読み込みエラー(t *testing.T) {
	// モック準備
	mockCSV := new(MockQry仕訳)
	mock仕訳XLSX := new(MockRep仕訳)
	mock勘定科目XLSX := new(MockQry勘定科目)

	// エラーを返すように設定
	readErr := errors.New("CSV読み込みエラー")
	mockCSV.On("ReadAll").Return(nil, readErr)

	// テスト対象
	service := NewService仕訳(mockCSV, mock仕訳XLSX, mock勘定科目XLSX)
	result, err := service.Execute仕訳集計()

	// 検証
	assert.Error(t, err)
	assert.True(t, errors.Is(err, Error仕訳読込失敗))
	// 失敗した場合は空のList集計仕訳が返ることを確認
	assert.Equal(t, NewList集計仕訳(), result)

	// モックの呼び出し検証
	mockCSV.AssertExpectations(t)
}

func TestQuery仕訳一覧_新規仕訳(t *testing.T) {
	// モック準備
	mockCSV := new(MockQry仕訳)
	mock仕訳XLSX := new(MockRep仕訳)
	mock勘定科目XLSX := new(MockQry勘定科目)

	// CSVの仕訳データ（新規）
	csv仕訳 := []*Ent仕訳{
		createSample仕訳(),
	}

	// XLSXの仕訳データ（空）
	xlsx仕訳 := []*Ent仕訳{}

	// 勘定科目データ
	勘定科目一覧 := []*Ent勘定科目{
		createSample勘定科目("旅費交通費", 基本ルール_経費配賦),
	}

	// モックの動作設定
	mockCSV.On("ReadAll").Return(csv仕訳, nil)
	mock仕訳XLSX.On("Read仕訳一覧").Return(xlsx仕訳, nil)
	mock勘定科目XLSX.On("Read勘定科目一覧").Return(勘定科目一覧, nil)

	// テスト対象
	service := NewService仕訳(mockCSV, mock仕訳XLSX, mock勘定科目XLSX)
	result, err := service.query仕訳一覧()

	// 検証
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "旅費交通費", result[0].Fld借方勘定科目)
	assert.NotNil(t, result[0].Val仕訳詳細)
	assert.Equal(t, "202405", result[0].Val仕訳詳細.Fld計上年月)
	assert.Equal(t, string(基本ルール_経費配賦), result[0].Val仕訳詳細.Fld按分ルール1)
	assert.Equal(t, "202405", result[0].Val仕訳詳細.Fld按分ルール2)

	// モックの呼び出し検証
	mockCSV.AssertExpectations(t)
	mock仕訳XLSX.AssertExpectations(t)
	mock勘定科目XLSX.AssertExpectations(t)
}

func TestQuery仕訳一覧_未定義勘定科目(t *testing.T) {
	// モック準備
	mockCSV := new(MockQry仕訳)
	mock仕訳XLSX := new(MockRep仕訳)
	mock勘定科目XLSX := new(MockQry勘定科目)

	// CSVの仕訳データ（存在しない勘定科目）
	csv仕訳 := []*Ent仕訳{
		func() *Ent仕訳 {
			j := createSample仕訳()
			j.Fld借方勘定科目 = "存在しない科目"
			return j
		}(),
	}

	// XLSXの仕訳データ（空）
	xlsx仕訳 := []*Ent仕訳{}

	// 勘定科目データ（対象科目なし）
	勘定科目一覧 := []*Ent勘定科目{
		createSample勘定科目("旅費交通費", 基本ルール_経費配賦),
	}

	// モックの動作設定
	mockCSV.On("ReadAll").Return(csv仕訳, nil)
	mock仕訳XLSX.On("Read仕訳一覧").Return(xlsx仕訳, nil)
	mock勘定科目XLSX.On("Read勘定科目一覧").Return(勘定科目一覧, nil)

	// テスト対象
	service := NewService仕訳(mockCSV, mock仕訳XLSX, mock勘定科目XLSX)
	result, err := service.query仕訳一覧()

	// 検証
	assert.Error(t, err)
	assert.True(t, errors.Is(err, Error未定義仕訳))
	assert.Len(t, result, 1)
	assert.Equal(t, "存在しない科目", result[0].Fld借方勘定科目)
	assert.Nil(t, result[0].Val仕訳詳細)

	// モックの呼び出し検証
	mockCSV.AssertExpectations(t)
	mock仕訳XLSX.AssertExpectations(t)
	mock勘定科目XLSX.AssertExpectations(t)
}

func TestSave(t *testing.T) {
	// モック準備
	mockCSV := new(MockQry仕訳)
	mock仕訳XLSX := new(MockRep仕訳)
	mock勘定科目XLSX := new(MockQry勘定科目)

	// 保存用の仕訳一覧
	仕訳一覧 := []*Ent仕訳{
		createSample仕訳(),
	}

	// モックの動作設定
	mock仕訳XLSX.On("Save", 仕訳一覧).Return(nil)

	// テスト対象
	service := NewService仕訳(mockCSV, mock仕訳XLSX, mock勘定科目XLSX)
	err := service.save(仕訳一覧)

	// 検証
	assert.NoError(t, err)

	// モックの呼び出し検証
	mock仕訳XLSX.AssertExpectations(t)
}

func TestQuery集計仕訳(t *testing.T) {
	// モック準備
	mockCSV := new(MockQry仕訳)
	mock仕訳XLSX := new(MockRep仕訳)
	mock勘定科目XLSX := new(MockQry勘定科目)

	// 仕訳一覧データ
	仕訳一覧 := []*Ent仕訳{
		func() *Ent仕訳 {
			j := createSample仕訳()
			j.Val仕訳詳細 = &Val仕訳詳細{
				Fld計上年月:   "202405",
				Fld原価要素:   "経費",
				Fldコストプール: "営業部",
				Fld按分ルール1: string(基本ルール_経費配賦),
				Fld按分ルール2: "202405",
			}
			return j
		}(),
		func() *Ent仕訳 {
			j := createSample仕訳()
			j.Val仕訳詳細 = &Val仕訳詳細{
				Fld計上年月:   "202405",
				Fld原価要素:   "経費",
				Fldコストプール: "営業部",
				Fld按分ルール1: string(基本ルール_経費配賦),
				Fld按分ルール2: "202405",
			}
			return j
		}(),
		func() *Ent仕訳 {
			j := createSample仕訳()
			j.FldNo = 2
			j.Fld借方金額 = decimal.NewFromInt(5000)
			j.Val仕訳詳細 = &Val仕訳詳細{
				Fld計上年月:   "202405",
				Fld原価要素:   "経費",
				Fldコストプール: "営業部",
				Fld按分ルール1: "対象外",
				Fld按分ルール2: "",
			}
			return j
		}(),
	}

	// テスト対象
	service := NewService仕訳(mockCSV, mock仕訳XLSX, mock勘定科目XLSX)
	result, err := service.query集計仕訳(仕訳一覧)

	// 検証
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// 「対象外」は集計しないため、1件のみ集計される
	集計仕訳一覧 := result.Get()
	assert.Len(t, 集計仕訳一覧, 1)

	// 集計結果の検証
	集計仕訳 := 集計仕訳一覧[0]
	assert.Equal(t, "202405", 集計仕訳.Fld計上年月)
	assert.Equal(t, "営業部", 集計仕訳.Fldコストプール)
	assert.Equal(t, "経費配賦", 集計仕訳.Fld按分ルール1)
	assert.Equal(t, "202405", 集計仕訳.Fld按分ルール2)
	assert.Equal(t, decimal.NewFromInt(20000), 集計仕訳.Fld合計金額)
}
