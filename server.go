package goe

import (
	"net/http"
	"strconv"
)

func InitServer(port int) {
	// TODO: Read config file
	initServerConfig()
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

	// TODO: Start server
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
