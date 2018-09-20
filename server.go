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
	RegStatus(http.StatusNotFound, func() string {
		return string(http.StatusNotFound)
	})

	// TODO: Exec runtime function
	queue := []string{runtimeStatus, runtimeListen, runtimeAction, runtimeRelate}
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(registerRuntime[queue[i]]); j++ {
			registerRuntime[queue[i]][j]()
		}
	}

	// TODO: Register to handle
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if action, ok := actionRegistry[request.URL.Path]; ok {
			// TODO: Check listen
			if request.Method == action.Method {
				for _, listen := range action.Listens {
					result := listen.Execute(writer, request)
					if result != nil {
						result(writer, request)
						break
					}
				}
				// TODO: Execute action
				action.Execute(writer, request)
				return
			}
		}
		httpState(writer, http.StatusNotFound)
	})

	// TODO: Register goe apis
	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "info")
	})

	// TODO: Start server
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func httpState(writer http.ResponseWriter, code int) {
	if f, ok := statusRegistry[code]; ok {
		if f != nil {
			http.Error(writer, f(), code)
			return
		}
	}
	httpState(writer, http.StatusNotFound)
}
