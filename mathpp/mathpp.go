package mathpp

import "math"

type (
	// 基础数学模型
	Model interface {
		// 可计算数字
		Number() float64
		// 模型可视化
		String() string
	}

	mathPP struct {
	}
)

var (
	O mathPP
)

// N进制下的极限加法表
func (mpp mathPP) TableP(rule uint8) *matrix {
	wh := int(rule)
	return mpp.CreateMatrix(wh, wh, func(j int, k int) Model {
		return &number{value: float64(j + k), rule: rule}
	})
}

// N进制下的极限乘法表
func (mpp mathPP) TableM(rule uint8) *matrix {
	wh := int(rule)
	return mpp.CreateMatrix(wh, wh, func(j int, k int) Model {
		return &number{value: float64(j * k), rule: rule}
	})
}

// 等差数列求和
func (mpp mathPP) ArithmeticProgression(a float64, d float64, n uint) float64 {
	return (float64(n) * (2*a + float64(n-1)*d)) / float64(2)
}

// 等比数列求和
func (mpp mathPP) GeometricProgression(a float64, q float64, n uint) float64 {
	return a * ((1 - math.Pow(q, float64(n))) / (1 - q))
}

// 前n项平方和
func (mpp mathPP) ItemsM2SumN(n uint) uint64 {
	return uint64((n * (n + 1) * (2*n + 1)) / 6)
}

// 前n项立方和
func (mpp mathPP) ItemsM3SumN(n uint) uint64 {
	return uint64(math.Pow(float64(n*(n+1)/2), 2))
}

// 二项式定理
func (mpp mathPP) BinomialTheorem() {

}
