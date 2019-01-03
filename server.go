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
	// 服务结构
	Server struct {
		mark    string            `json:"mark"` // 硬件唯一标识符
		Name    string            `json:"name"` // 服务名,比如说dmbr,stio之类的
		host    string            `json:"-"`    // 当前Host
		Port    int               `json:"port"` // 当前Port
		child   []Server          `json:"-"`    // 子集服务
		parent  *Server           `json:"-"`    // 上级服务
		servers []Server          `json:"-"`    // 兄弟服务
		vendors map[string]Server `json:"-"`    // 供应商服务
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

func (s Server) InitServer() {
	// 初始化当前服务
	host, _ := os.Hostname()
	currentServer = Server{
		mark:    uniqueMark(),
		Name:    s.Name,
		host:    host,
		Port:    s.Port,
		servers: []Server{},
		vendors: make(map[string]Server),
	}
	// 强制更新冰箱配置文件
	iceConfig.update()
	//if len(os.Args) == 1 {
	//} else {
	//	// 检查 args[1]参数是否符合host:port
	//	// 从host:port获取数据
	//}
	// 注册默认的状态码展示页
	for i := 0; i < len(httpStatus); i++ {
		s.RegStatus(httpStatus[i], func(code int) string {
			return strconv.Itoa(code) + " " + http.StatusText(code)
		})
	}
	// 注册 GOE 系统级API
	// initSystemApis()
	// 执行注册运行时,把自定义状态码,监听器,控制器,关联关系(监听器&控制器)
	queue := []string{manageStatus, manageListen, manageAction, manageRelate}
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(serverManage[queue[i]]); j++ {
			serverManage[queue[i]][j]()
		}
	}
	// 启动GOE总路由监听
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		out := Out{w: writer}
		if action, ok := actionRegistry[request.URL.Path]; ok {
			in := In{r: request}
			// TODO: Check listen
			if action.methodContains(request.Method) {
				for _, listen := range action.Listens {
					result := listen.Process(in)
					if result != nil {
						result(in, out)
						return
					}
				}
				// TODO: Exec action
				action.Program(in, out)
				return
			}
		}
		out.status(http.StatusNotFound)
	})
	// TODO: Start server
	_ = http.ListenAndServe(":"+strconv.Itoa(currentServer.Port), nil)
}

func (s Server) RegStatus(code int, f func(int) string) {
	if f != nil {
		joinManage(manageStatus, func() { statusRegistry[code] = f })
	}
}
