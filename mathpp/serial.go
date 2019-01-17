package mathpp

import "math"

type (
	// 序列结构
	serial []Model
)

// 创建一个定长的序列,并逐项构建
func (mathPP) CreateSerial(w int, item func(j int) Model) *serial {
	serial := &serial{}
	for j := 0; j < w; j++ {
		*serial = append(*serial, item(j))
	}
	return serial
}

// 模型可视化
func (s *serial) String() string {
	result := ""
	for _, r := range *s {
		result += r.String() + "\t"
	}
	return result
}

// 序列总和
func (s *serial) Sum() *number {
	sum := &number{rule: 10}
	for _, n := range *s {
		sum.value += n.Number()
	}
	return sum
}

// 平均数
func (s *serial) Avg() *number {
	return &number{
		rule:  10,
		value: s.Sum().Number() / float64(len(*s)),
	}
}

// 中位数
func (s *serial) Med() *number {
	med := &number{rule: 10}
	cnt := len(*s)
	haf := math.Floor(float64(cnt) / 2)
	if cnt%2 == 1 {
		med.value = (*s)[int(haf)].Number()
	} else {
		med.value = ((*s)[int(haf)].Number() + (*s)[int(haf+1)].Number()) / 2
	}
	return med
}
