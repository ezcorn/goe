package data

type (
	// 模型手柄
	shell struct {
		model     *Model
		condition []string
	}
)

func Shell(model Model) *shell {
	return &shell{
		model: &model,
	}
}

func (s *shell) Where(attr string, args ...interface{}) *shell {
	return s
}

func (s *shell) Count() uint {
	return 0
}
