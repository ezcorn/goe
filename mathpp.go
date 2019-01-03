package goe

import (
	"math"
	"strings"
)

type (
	mathPP struct {
	}
	// 数字
	n struct {
		numb float64 // 十进制数
		rule uint8   // 进位制度
	}

	// 单行
	r []*n

	// 矩阵
	m []*r
)

var (
	MathPP mathPP
	// 进位制度符号
	ruleSymbol = [][2]byte{{'0', '9'}, {'a', 'z'}, {'A', 'Z'}}
)

/**
 * 生成一个十进制数字
 */
func (mathPP) N(numb float64) *n {
	i := &n{numb: numb, rule: 10}
	return i
}

/**
 * 生成一个N进制数字
 */
func (mathPP) NRule(numb float64, rule uint8) *n {
	// 不支持0进制和1进制
	if rule < 2 {
		return nil
	}
	i := &n{numb: numb, rule: rule}
	return i
}

/**
 * 通过数字展示和进位制度生成一个数字
 */
func (mathPP) NStr(view string, rule uint8) *n {
	// 不支持0进制和1进制
	if rule < 2 {
		return nil
	}
	numb := float64(0)
	view = strings.TrimSpace(view)
	for k, char := range view {
		i := 0
		for _, rge := range ruleSymbol {
			for j := rge[0]; j <= rge[1]; j++ {
				if byte(char) == j {
					goto BREAK
				}
				i++
			}
		}
		return nil
	BREAK:
		numb += float64(i) * math.Pow(float64(rule), float64(k))
	}
	return &n{numb: numb, rule: rule}
}

/**
 ******************** N 系列 ********************
 */

/**
 * 生成一个数字的显示模式
 */
func (i *n) String() string {
	view := ""
	if i == nil {
		return view
	}
	result := i.numb
	for {
		rem := int64(result) % int64(i.rule)
		for _, rge := range ruleSymbol {
			for j := rge[0]; j <= rge[1]; j++ {
				if rem == 0 {
					view = string(byte(j)) + view
					goto BREAK
				}
				rem--
			}
		}
	BREAK:
		result = float64(math.Floor(float64(result) / float64(i.rule)))
		if result == 0 {
			break
		}
	}
	return view
}

/**
 * 重新设置这个数字的进位机制
 */
func (i *n) Rule(rule uint8) *n {
	if i == nil {
		return nil
	}
	i.rule = rule
	return i
}

/**
 * 前置运算检查
 */
func (i *n) preOperate(i2 *n, fun func() *n) *n {
	if i == nil || i2 == nil {
		return nil
	}
	return fun()
}

/**
 * 数字加法
 */
func (i *n) Plus(i2 *n) *n {
	return i.preOperate(i2, func() *n {
		return &n{numb: i.numb + i2.numb, rule: i.rule}
	})
}

/**
 * 数字减法
 */
func (i *n) Less(i2 *n) *n {
	return i.preOperate(i2, func() *n {
		return &n{numb: i.numb - i2.numb, rule: i.rule}
	})
}

/**
 * 数字乘法
 */
func (i *n) Multi(i2 *n) *n {
	return i.preOperate(i2, func() *n {
		return &n{numb: i.numb * i2.numb, rule: i.rule}
	})
}

/**
 * 数字除法
 */
func (i *n) Division(i2 *n) *n {
	return i.preOperate(i2, func() *n {
		return &n{numb: i.numb / i2.numb, rule: i.rule}
	})
}

/**
 * 数字模运算
 */
func (i *n) Mod(i2 *n) *n {
	return i.preOperate(i2, func() *n {
		return &n{numb: float64(int64(i.numb) % int64(i2.numb)), rule: i.rule}
	})
}

/**
 * 加法表
 */
func (mpp mathPP) PlusTable(rule uint8) *m {
	wh := int(rule)
	return mpp.M(wh, wh, func(j int, k int) *n {
		return &n{numb: float64(j + k), rule: rule}
	})
}

/**
 * 乘法表
 */
func (mpp mathPP) MultiTable(rule uint8) *m {
	wh := int(rule)
	return mpp.M(wh, wh, func(j int, k int) *n {
		return &n{numb: float64(j * k), rule: rule}
	})
}

/**
 ******************** M 系列 ********************
 */

func (mathPP) M(w int, h int, item func(j int, k int) *n) *m {
	m := &m{}
	var j, k int
	for j = 0; j < h; j++ {
		r := &r{}
		for k = 0; k < w; k++ {
			*r = append(*r, item(j, k))
		}
		*m = append(*m, r)
	}
	return m
}

func (r *r) String() string {
	result := ""
	for _, v := range *r {
		result += v.String() + "\t"
	}
	return result
}

func (m *m) String() string {
	result := ""
	for _, r := range *m {
		result += r.String() + "\n"
	}
	return result
}
