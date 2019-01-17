package goe

import "encoding/json"

type (
	jsn struct{}
)

var (
	JSON jsn
)

// JSON反序列化
func (jsn) Decode(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		v = nil
	}
}

// JSON序列号
func (jsn) Encode(data interface{}) string {
	j, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(j)
}

// JSON序列格式化
func (jsn) EncodeIndent(data interface{}) string {
	j, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return ""
	}
	return string(j)
}
