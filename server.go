package goe

import (
	"fmt"
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
