package mpp

import (
	"math"
	"strings"
)

type (
	// 数模型
	number struct {
		value float64 // 十进制数
		rule  uint8   // 进位制度
	}
)

var (
	// 非十进制情况下的符号展示集
	numberRuleSymbol = [][2]byte{{'0', '9'}, {'a', 'z'}, {'A', 'Z'}}
)

// 基于数值生成一个十进制数
func (mpp mathPP) CreateNumber(value float64) *number {
	return mpp.CreateNumberFromRule(value, 10)
}

// 基于数值和进制生成一个数
func (mathPP) CreateNumberFromRule(value float64, rule uint8) *number {
	return &number{value: value, rule: rule}
}

// 基于字符串和进制还原一个数
func (mathPP) CreateNumberFromString(str string, rule uint8) *number {
	// 不支持0进制和1进制
	if rule < 2 {
		return nil
	}
	value := float64(0)
	str = strings.TrimSpace(str)
	for k, char := range str {
		i := 0
		for _, rge := range numberRuleSymbol {
			for j := rge[0]; j <= rge[1]; j++ {
				if byte(char) == j {
					goto BREAK
				}
				i++
			}
		}
		return nil
	BREAK:
		value += float64(i) * math.Pow(float64(rule), float64(k))
	}
	return &number{value: value, rule: rule}
}

// 可计算数字
func (n *number) Number() float64 {
	return n.value
}

// 模型可视化
func (n *number) String() string {
	view := ""
	if n == nil {
		return view
	}
	if math.IsNaN(n.Number()) {
		return "NaN"
	}
	result := n.Number()
	for {
		rem := int64(result) % int64(n.rule)
		for _, rge := range numberRuleSymbol {
			for j := rge[0]; j <= rge[1]; j++ {
				if rem == 0 {
					view = string(byte(j)) + view
					goto BREAK
				}
				rem--
			}
		}
	BREAK:
		result = float64(math.Floor(float64(result) / float64(n.rule)))
		if result == 0 {
			break
		}
	}
	return view
}

// 加法
func (n *number) P(n1 *number) *number {
	return &number{value: n.Number() + n1.Number(), rule: n.rule}
}

// 减法
func (n *number) S(n1 *number) *number {
	return &number{value: n.Number() - n1.Number(), rule: n.rule}
}

// 乘法
func (n *number) M(n1 *number) *number {
	return &number{value: n.Number() * n1.Number(), rule: n.rule}
}

// 除法
func (n *number) D(n1 *number) *number {
	return &number{value: n.Number() / n1.Number(), rule: n.rule}
}

// 求余
func (n *number) R(n1 *number) *number {
	return &number{value: float64(int64(n.Number()) % int64(n1.Number())), rule: n.rule}
}
