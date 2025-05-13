package domain

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestEnt按分結果_Calc税(t *testing.T) {
	tests := []struct {
		name     string
		entity   Ent按分結果
		expected decimal.Decimal
	}{
		{
			name: "税率が0の場合",
			entity: Ent按分結果{
				Fld金額:   decimal.NewFromInt(1000),
				Fld借方税率: decimal.Zero,
			},
			expected: decimal.Zero,
		},
		{
			name: "借方税区分が控80の場合",
			entity: Ent按分結果{
				Fld金額:    decimal.NewFromInt(206542),
				Fld借方税率:  decimal.NewFromInt(10),
				Fld借方税区分: "控80",
			},
			expected: decimal.NewFromInt(16199), // 手取りが200,000 円になるように、222,741 円払った場合の例
		},
		{
			name: "通常の税率計算の場合",
			entity: Ent按分結果{
				Fld金額:    decimal.NewFromInt(1000),
				Fld借方税率:  decimal.NewFromInt(10),
				Fld借方税区分: "課税",
			},
			expected: decimal.NewFromInt(100), // 1000 * 10 / 100 = 100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.entity.Calc税()
			if !tt.expected.Equal(actual) {
				t.Errorf("期待値: %s, 実際: %s", tt.expected.String(), actual.String())
			}
		})
	}
}

func TestEnt按分結果_Calc税込金額(t *testing.T) {
	tests := []struct {
		name     string
		entity   Ent按分結果
		expected decimal.Decimal
	}{
		{
			name: "税率が0の場合",
			entity: Ent按分結果{
				Fld金額:   decimal.NewFromInt(1000),
				Fld借方税率: decimal.Zero,
			},
			expected: decimal.NewFromInt(1000),
		},
		{
			name: "借方税区分が控80の場合",
			entity: Ent按分結果{
				Fld金額:    decimal.NewFromInt(206542),
				Fld借方税率:  decimal.NewFromInt(10),
				Fld借方税区分: "控80",
			},
			expected: decimal.NewFromInt(222741), // 206542 + 16199
		},
		{
			name: "通常の税率計算の場合",
			entity: Ent按分結果{
				Fld金額:    decimal.NewFromInt(1000),
				Fld借方税率:  decimal.NewFromInt(10),
				Fld借方税区分: "課税",
			},
			expected: decimal.NewFromInt(1100), // 1000 + 100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.entity.Calc税込金額()
			if !tt.expected.Equal(actual) {
				t.Errorf("期待値: %s, 実際: %s", tt.expected.String(), actual.String())
			}
		})
	}
}
