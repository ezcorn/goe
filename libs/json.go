package libs

import "encoding/json"

type (
	Json struct{}
)

func (Json) Encode(data interface{}) string {
	j, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(j)
}

func (Json) Decode(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		v = nil
	}
}
