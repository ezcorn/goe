package goe

import (
	"net/http"
	"strconv"
)

type (
	Server struct {
		Name    string             `json:"name"` // 服务名
		Host    string             `json:"host"` // 当前Host
		Port    int                `json:"port"` // 当前Port
		Servers []*Server          `json:"-"`    // 兄弟服务
		Vendors map[string]*Server `json:"-"`    // 供应商服务
	}
)

var (
	currentServer = Server{}
)

func InitServer(port int) {
	// TODO: Register default state print
	initServerStatus()
	// TODO: Exec runtime
	initServerManage()
	// TODO: Goe framework
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		out := Out{w: writer}
		if action, ok := actionRegistry[request.URL.Path]; ok {
			in := In{r: request}
			// TODO: Check listen
			if action.MethodContains(request.Method) {
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
		out.Status(http.StatusNotFound)
	})

	// TODO: Register goe apis
	initServerSentry()

	// TODO: Start vendor task
	initVendorTask()

	// TODO: Start server
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
