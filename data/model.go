package data

type (
	// 基础数学模型
	Model interface {
		// 可计算数字
		Number() float64
		// 模型可视化
		String() string
		// 对应关系
		Relate() map[string]string
		// 是否开启缓存
		UseCache() bool
	}
)
