package domain

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// AllocationResult は元の要素と按分結果を保持する構造体です。
// ジェネリック型 T を使用して、任意の型の元の要素を保持できます。
type AllocationResult[T any] struct {
	Original       T               // 元のリストの要素
	AllocatedValue decimal.Decimal // 按分後の値
	DiffValue      decimal.Decimal // 按分誤差
}

type roundingMode int

const (
	RoundHalfUp roundingMode = iota
	RoundDown
	RoundUp
)

// allocateConfig は Calc按分 関数の内部設定を保持します。
type allocateConfig struct {
	scale        int32        // 計算結果の小数点以下の桁数
	roundingMode roundingMode // 丸め方法
}

// OptionFunc は allocateConfig を変更するための関数型です。
type OptionFunc func(*allocateConfig)

// WithScale は按分計算の小数点以下の桁数を指定するオプションです。
// scale は 0 以上である必要があります。
// デフォルト: 引数 total の小数点以下の桁数。
func WithScale(scale int32) OptionFunc {
	return func(c *allocateConfig) {
		if scale < 0 {
			fmt.Println("Warning: WithScale received negative value, using 0 instead.")
			scale = 0
		}
		c.scale = scale
	}
}

// WithRoundingMode は按分計算時の丸めモードを指定するオプションです。
// 利用可能なモードは RoundHalfUp (四捨五入), RoundDown (切り捨て), RoundUp (切り上げ) です。
// デフォルト: RoundHalfUp (四捨五入)。
func WithRoundingMode(mode roundingMode) OptionFunc {
	return func(c *allocateConfig) {
		c.roundingMode = mode
	}
}

// Calc按分 は按分計算の関数です。
func Calc按分[T any](
	total decimal.Decimal,
	items []T,
	getWeight func(e T) decimal.Decimal,
	opts ...OptionFunc,
) ([]AllocationResult[T], error) {
	if len(items) == 0 {
		return nil, errors.New("items cannot be empty")
	}
	// デフォルト設定
	config := allocateConfig{
		scale:        0,
		roundingMode: RoundHalfUp, // 四捨五入
	}

	// 提供されたオプションを適用
	for _, opt := range opts {
		opt(&config)
	}

	w := make([]decimal.Decimal, len(items))
	sumW := decimal.Zero
	tmpMax := decimal.Zero
	maxI := 0
	for i, item := range items {
		w[i] = getWeight(item)
		sumW = sumW.Add(w[i])
		if tmpMax.Cmp(w[i]) == -1 {
			tmpMax = w[i]
			maxI = i
		}
	}
	if sumW.IsZero() {
		return nil, errors.New("sum of weights cannot be zero")
	}

	ret := make([]AllocationResult[T], len(w))
	sumR := decimal.Zero
	for i := range w {
		ret[i].Original = items[i]
		allocatedValue := total.Mul(w[i]).Div(sumW)
		switch config.roundingMode {
		case RoundHalfUp:
			allocatedValue = allocatedValue.Round(config.scale)
		case RoundDown:
			allocatedValue = allocatedValue.RoundDown(config.scale)
		case RoundUp:
			allocatedValue = allocatedValue.RoundUp(config.scale)
		}
		ret[i].AllocatedValue = allocatedValue
		sumR = sumR.Add(allocatedValue)
	}
	diff := total.Add(sumR.Neg())
	if !diff.IsZero() {
		ret[maxI].AllocatedValue = ret[maxI].AllocatedValue.Add(diff)
		ret[maxI].DiffValue = diff
	}
	return ret, nil
}
