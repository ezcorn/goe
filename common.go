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

func MakeServer(name string, port int) *Server {
	return &Server{
		Name: name,
		Port: port,
	}
}

func MakeAction(route string, comment string, method []string, program program) *Action {
	if program == nil {
		program = func(in In, out Out, libs Libs) {}
	}
	return &Action{
		Route:   route,
		Method:  method,
		Comment: comment,
		Listens: listens{},
		Program: program,
	}
}

func MakeListen(name string, comment string, process process) *listen {
	if process == nil {
		process = func(in In, libs Libs) program { return nil }
	}
	return &listen{
		Name:    name,
		Comment: comment,
		Process: process,
	}
}
