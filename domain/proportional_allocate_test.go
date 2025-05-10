package domain

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestProportionalAllocate(t *testing.T) {
	type testItem struct {
		Name   string
		Weight decimal.Decimal
	}

	type testCase struct {
		name                 string
		total                decimal.Decimal
		items                []testItem
		scale                int32
		roundingMode         roundingMode
		expectedAllocated    []decimal.Decimal
		expectedDiff         []decimal.Decimal
		expectedDiffItemName string // 誤差調整が適用される要素の名前
	}

	tests := []testCase{
		{
			name:         "基本的な按分計算",
			total:        decimal.NewFromInt(600),
			items:        []testItem{{"A", decimal.NewFromInt(1)}, {"B", decimal.NewFromInt(2)}, {"C", decimal.NewFromInt(3)}},
			scale:        0,
			roundingMode: RoundHalfUp,
			expectedAllocated: []decimal.Decimal{
				decimal.NewFromInt(100),
				decimal.NewFromInt(200),
				decimal.NewFromInt(300),
			},
			expectedDiff: []decimal.Decimal{
				decimal.Zero,
				decimal.Zero,
				decimal.Zero,
			},
		},
		{
			name:         "RoundHalfUpでの丸め誤差",
			total:        decimal.NewFromInt(100),
			items:        []testItem{{"A", decimal.NewFromInt(1)}, {"B", decimal.NewFromInt(1)}, {"C", decimal.NewFromInt(1)}},
			scale:        0,
			roundingMode: RoundHalfUp,
			expectedAllocated: []decimal.Decimal{
				decimal.NewFromInt(34),
				decimal.NewFromInt(33),
				decimal.NewFromInt(33),
			},
			expectedDiff: []decimal.Decimal{
				decimal.NewFromInt(1),
				decimal.Zero,
				decimal.Zero,
			},
			expectedDiffItemName: "A",
		},
		{
			name:         "RoundDownでの丸め誤差",
			total:        decimal.NewFromInt(100),
			items:        []testItem{{"A", decimal.NewFromInt(1)}, {"B", decimal.NewFromInt(1)}, {"C", decimal.NewFromInt(1)}},
			scale:        0,
			roundingMode: RoundDown,
			expectedAllocated: []decimal.Decimal{
				decimal.NewFromInt(34),
				decimal.NewFromInt(33),
				decimal.NewFromInt(33),
			},
			expectedDiff: []decimal.Decimal{
				decimal.NewFromInt(1),
				decimal.Zero,
				decimal.Zero,
			},
			expectedDiffItemName: "A",
		},
		{
			name:         "RoundUpでの丸め誤差",
			total:        decimal.NewFromInt(100),
			items:        []testItem{{"A", decimal.NewFromInt(1)}, {"B", decimal.NewFromInt(1)}, {"C", decimal.NewFromInt(1)}},
			scale:        0,
			roundingMode: RoundUp,
			expectedAllocated: []decimal.Decimal{
				decimal.NewFromInt(32),
				decimal.NewFromInt(34),
				decimal.NewFromInt(34),
			},
			expectedDiff: []decimal.Decimal{
				decimal.NewFromInt(-2),
				decimal.Zero,
				decimal.Zero,
			},
			expectedDiffItemName: "A",
		},
		{
			name:         "異なる重みでの丸め誤差",
			total:        decimal.NewFromInt(100),
			items:        []testItem{{"A", decimal.NewFromInt(10)}, {"B", decimal.NewFromInt(20)}, {"C", decimal.NewFromInt(15)}},
			scale:        0,
			roundingMode: RoundHalfUp,
			expectedAllocated: []decimal.Decimal{
				decimal.NewFromInt(22),
				decimal.NewFromInt(45),
				decimal.NewFromInt(33),
			},
			expectedDiff: []decimal.Decimal{
				decimal.Zero,
				decimal.NewFromInt(1),
				decimal.Zero,
			},
			expectedDiffItemName: "B",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			getWeight := func(e testItem) decimal.Decimal {
				return e.Weight
			}

			results, err := ProportionalAllocate(
				tc.total,
				tc.items,
				getWeight,
				WithScale(tc.scale),
				WithRoundingMode(tc.roundingMode),
			)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if len(results) != len(tc.items) {
				t.Fatalf("expected %d results, got %d", len(tc.items), len(results))
			}

			// 結果をマップに変換して名前で検索できるようにする
			resultMap := make(map[string]AllocationResult[testItem])
			for _, res := range results {
				resultMap[res.Original.Name] = res
			}

			// 期待値の検証
			for i, item := range tc.items {
				res := resultMap[item.Name]
				if !tc.expectedAllocated[i].Equal(res.AllocatedValue) {
					t.Errorf("Item %s AllocatedValue mismatch: expected %v, got %v",
						item.Name, tc.expectedAllocated[i], res.AllocatedValue)
				}
				if !tc.expectedDiff[i].Equal(res.DiffValue) {
					t.Errorf("Item %s DiffValue mismatch: expected %v, got %v",
						item.Name, tc.expectedDiff[i], res.DiffValue)
				}

				// 誤差調整が期待される要素の確認
				if tc.expectedDiffItemName != "" && item.Name == tc.expectedDiffItemName {
					if res.DiffValue.IsZero() {
						t.Errorf("Expected diff adjustment on item %s, but DiffValue is zero", item.Name)
					}
				}
			}
		})
	}
}

// エラーケースのテスト
func TestProportionalAllocateErrors(t *testing.T) {
	tests := []struct {
		name          string
		total         decimal.Decimal
		items         interface{}
		getWeight     interface{}
		expectedError string
	}{
		{
			name:  "空の配列",
			total: decimal.NewFromInt(100),
			items: []struct{}{},
			getWeight: func(e struct{}) decimal.Decimal {
				return decimal.NewFromInt(1)
			},
			expectedError: "items cannot be empty",
		},
		{
			name:  "重みの合計が0",
			total: decimal.NewFromInt(100),
			items: []struct {
				Name   string
				Weight decimal.Decimal
			}{
				{"A", decimal.Zero},
				{"B", decimal.Zero},
			},
			getWeight: func(e struct {
				Name   string
				Weight decimal.Decimal
			}) decimal.Decimal {
				return e.Weight
			},
			expectedError: "sum of weights cannot be zero",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch items := tc.items.(type) {
			case []struct{}:
				getWeight := tc.getWeight.(func(struct{}) decimal.Decimal)
				_, err := ProportionalAllocate(tc.total, items, getWeight)
				if err == nil {
					t.Fatal("expected an error, got none")
				}
				if err.Error() != tc.expectedError {
					t.Fatalf("expected error message '%s', got '%v'", tc.expectedError, err.Error())
				}
			case []struct {
				Name   string
				Weight decimal.Decimal
			}:
				getWeight := tc.getWeight.(func(struct {
					Name   string
					Weight decimal.Decimal
				}) decimal.Decimal)
				_, err := ProportionalAllocate(tc.total, items, getWeight)
				if err == nil {
					t.Fatal("expected an error, got none")
				}
				if err.Error() != tc.expectedError {
					t.Fatalf("expected error message '%s', got '%v'", tc.expectedError, err.Error())
				}
			}
		})
	}
}
