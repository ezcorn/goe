package goe

import (
	"encoding/json"
)

type (
	// 标准字典重定义
	Map map[string]interface{}
	// 标准数组重定义
	Arr []interface{}
)

func jsonEncode(data interface{}) string {
	j, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(j)
}

func jsonDecode(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		v = nil
	}
}

func joinManage(t string, f func()) {
	serverManage[t] = append(serverManage[t], f)
}
