package goe

import (
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

// 唯一哈希(随机)
func Unique(size uint8) string {
	return Crypto.MD5(strconv.FormatInt(time.Now().UnixNano()+rand.Int63(), 10), size)
}

// 输出分组日志
func LogPrintln(group string, info string) {
	log.Println("[ " + group + " ] : " + info)
}

// 生成一个行为
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

// 生成一个监听器
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

// 生成一个服务
func MakeServer(name string, port int) *Server {
	return &Server{
		Name: name,
		Port: port,
	}
}
