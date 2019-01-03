package goe

import (
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type (
	// 标准字典重定义
	Map map[string]interface{}
	// 标准数组重定义
	Arr []interface{}
)

func jsonDecode(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		v = nil
	}
}

func jsonEncode(data interface{}) string {
	j, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(j)
}

func jsonEncodeIndent(data interface{}) string {
	j, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return ""
	}
	return string(j)
}

func joinManage(t string, f func()) {
	serverManage[t] = append(serverManage[t], f)
}

func uniqueMark() string {
	return Crypto.MD5(jsonEncode(Device.Network.Mac()))
}

func uniqueHash() string {
	return Crypto.MD5(strconv.FormatInt(time.Now().UnixNano()+rand.Int63(), 10))
}

func logPrintln(module string, info string) {
	log.Println("[ " + module + " ] : " + info)
}
func MakeServer(name string, port int) *Server {
	return &Server{
		Name: name,
		Port: port,
	}
}

func MakeAction(route string, comment string, method []string, program program) *Action {
	if program == nil {
		program = func(in In, out Out) {}
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
		process = func(in In) program { return nil }
	}
	return &listen{
		Name:    name,
		Comment: comment,
		Process: process,
	}
}
