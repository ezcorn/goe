package goe

import (
	"net/http"
	"os"
	"strconv"
)

const (
	manageStatus = "status"
	manageListen = "listen"
	manageAction = "action"
	manageRelate = "relate"
)

type (
	Server struct {
		Name    string            `json:"name"` // 服务名,比如说dmbr,stio之类的
		Host    string            `json:"host"` // 当前Host
		Port    int               `json:"port"` // 当前Port
		Child   []Server          `json:"-"`    // 子集服务
		Parent  *Server           `json:"-"`    // 上级服务
		Servers []Server          `json:"-"`    // 兄弟服务
		Vendors map[string]Server `json:"-"`    // 供应商服务
	}
)

var (
	currentServer  Server
	serverManage   = make(map[string][]func())
	statusRegistry = make(map[int]func(int) string)
	httpStatus     = []int{
		100, 101, 102,
		200, 201, 202, 203, 204, 205, 206, 207, 208, 226,
		300, 301, 302, 303, 304, 305, 307, 308,
		400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 426, 428, 429, 431, 451,
		500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511,
	}
)

func InitServer(name string, port int) {
	// TODO: Init system
	host, _ := os.Hostname()
	currentServer = Server{
		Name:    name,
		Host:    host,
		Port:    port,
		Servers: []Server{},
		Vendors: make(map[string]Server),
	}
	//if len(os.Args) == 1 {
	//} else {
	//	// 检查 args[1]参数是否符合host:port
	//	// 从host:port获取数据
	//}
	// TODO: Register default state print
	for i := 0; i < len(httpStatus); i++ {
		RegStatus(httpStatus[i], func(code int) string {
			return strconv.Itoa(code) + " " + http.StatusText(code)
		})
	}
	// TODO: Register goe apis
	// initSystemApis()
	// TODO: Exec runtime
	queue := []string{manageStatus, manageListen, manageAction, manageRelate}
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(serverManage[queue[i]]); j++ {
			serverManage[queue[i]][j]()
		}
	}

	// TODO: Init libs
	var libs Libs

	// TODO: Goe framework
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		out := Out{w: writer, libs: libs}
		if action, ok := actionRegistry[request.URL.Path]; ok {
			in := In{r: request, libs: libs}
			// TODO: Check listen
			if action.MethodContains(request.Method) {
				for _, listen := range action.Listens {
					result := listen.Process(in, libs)
					if result != nil {
						result(in, out, libs)
						return
					}
				}
				// TODO: Exec action
				action.Program(in, out, libs)
				return
			}
		}
		out.status(http.StatusNotFound)
	})
	// TODO: Start server
	_ = http.ListenAndServe(":"+strconv.Itoa(currentServer.Port), nil)
}

func joinManage(t string, f func()) {
	serverManage[t] = append(serverManage[t], f)
}

func RegStatus(code int, f func(int) string) {
	if f != nil {
		joinManage(manageStatus, func() { statusRegistry[code] = f })
	}
}
