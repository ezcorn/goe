package goe

import (
	"fmt"
	"net/http"
	"strconv"
)

func InitServer(port int) {
	// TODO: Read config file
	// json.Unmarshal(readFile("config.json"), &serverConfig)
	// initVendorTask(serverConfig.Vendor)

	// TODO: Register default state print
	for i := 0; i < len(httpStatus); i++ {
		RegStatus(httpStatus[i], func(code int) string {
			return strconv.Itoa(code) + " " + http.StatusText(code)
		})
	}

	// TODO: Exec runtime
	queue := []string{runtimeStatus, runtimeListen, runtimeAction, runtimeRelate}
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(registerRuntime[queue[i]]); j++ {
			registerRuntime[queue[i]][j]()
		}
	}

	// TODO: Goe framework
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if action, ok := actionRegistry[r.URL.Path]; ok {
			// TODO: Check listen
			if strArrContains(action.Method, r.Method) {
				for _, listen := range action.Listens {
					result := listen.Process(w, r)
					if result != nil {
						result(w, r)
						break
					}
				}
				// TODO: Exec action
				action.Program(w, r)
				return
			}
		}
		httpState(w, http.StatusNotFound)
	})

	// TODO: Register goe apis
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "info")
	})

	// TODO: Start server
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func httpState(w http.ResponseWriter, code int) {
	if f, ok := statusRegistry[code]; ok {
		if f != nil {
			http.Error(w, f(code), code)
			return
		}
	}
	httpState(w, http.StatusNotFound)
}

func Echo(w http.ResponseWriter, content string) {
	fmt.Fprintf(w, content)
}
