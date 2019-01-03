package goe

import (
	"fmt"
	"math"
)

type (
	mathPP struct {
	}
	// 数字结构
	integer struct {
		numb int64 // 十进制数
		base uint8 // 基于进制
	}
)

var (
	MathPP mathPP
	// 基于进制范围
	baseRange = [][2]byte{{'0', '9'}, {'a', 'z'}, {'A', 'Z'}}
)

/**
 * 通过数字展示和基于进制生成一个数字
 */
//func (math) IntegerFromBase(view string, base uint8) *integer {
//	// i := &integer{numb: numb, base: base}
//	// i.makeView()
//	// return i
//}

/**
 * 生成一个十进制数字
 */
func (m mathPP) Integer(numb int64) *integer {
	i := &integer{numb: numb, base: 10}
	return i
}

func (i *integer) View() string {
	view := ""
	result := i.numb
	for {
		rem := result % int64(i.base)
		for _, rge := range baseRange {
			for j := rge[0]; j < rge[1]; j++ {
				rem--
				if rem == 0 {
					fmt.Println(byte(j))
				}
			}
		}
		result = int64(math.Floor(float64(result) / float64(i.base)))
		if result == 0 {
			break
		}
	}
	return view
}

//func (i *integer) ViewBase() string {
//
//}

//func (i *integer) Change(base uint8) integer {
//	// i.numb
//}
