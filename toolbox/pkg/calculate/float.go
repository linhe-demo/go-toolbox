package calculate

import (
	"github.com/shopspring/decimal"
	"math"
)

const MIN = 0.000001

// AccuracyAdd 浮点数加
func AccuracyAdd(num1 float64, num2 float64) (out float64) {
	tmpVal := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	out, _ = tmpVal.Float64()
	return out
}

// AccuracyMin 浮点数减
func AccuracyMin(num1 float64, num2 float64) (out float64) {
	tmpVal := decimal.NewFromFloat(num1).Sub(decimal.NewFromFloat(num2))
	out, _ = tmpVal.Float64()
	return out
}

// AccuracyMul 浮点数乘
func AccuracyMul(num1 float64, num2 float64) (out float64) {
	tmpVal := decimal.NewFromFloat(num1).Mul(decimal.NewFromFloat(num2))
	out, _ = tmpVal.Float64()
	return out
}

// AccuracyDiv 浮点数除
func AccuracyDiv(num1 float64, num2 float64) (out float64) {
	tmpVal := decimal.NewFromFloat(num1).Div(decimal.NewFromFloat(num2))
	out, _ = tmpVal.Float64()
	return out
}

// AccuracyRound 浮点数保留几位有效数字
func AccuracyRound(target float64, digits int32) (out float64) {
	out, _ = decimal.NewFromFloat(target).Round(digits).Float64()
	return out
}

// IsEqual 判断 float 是否相等
func IsEqual(x, y float64) bool {
	return math.Abs(x-y) < MIN
}

// Wrap 将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}
