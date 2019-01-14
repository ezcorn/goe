package mathpp

type (
	// 矩阵结构
	matrix []*serial
)

// 创建一个定长高的矩阵,并逐项构建
func (mpp mathPP) CreateMatrix(w int, h int, item func(i int, j int) Model) *matrix {
	matrix := &matrix{}
	for i := 0; i < h; i++ {
		*matrix = append(*matrix, mpp.CreateSerial(w, func(j int) Model {
			return item(i, j)
		}))
	}
	return matrix
}

func (matrix *matrix) String() string {
	result := ""
	for _, v := range *matrix {
		result += v.String() + "\n"
	}
	return result
}
