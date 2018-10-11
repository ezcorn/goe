package libs

type (
	Memory struct{}
)

func (m Memory) LocalSet(key string, data interface{}) {
}

func (m Memory) LocalGet(key string) interface{} {
	return nil
}

func (m Memory) ShareSet(key string, data interface{}) {
}

func (m Memory) ShareGet(key string) interface{} {
	return nil
}
