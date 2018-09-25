package goe

import (
	"encoding/json"
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
			if arrContainsStr(action.Method, r.Method) {
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
	initServerSentry()

	// TODO: Start server
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func arrContainsStr(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func Echo(w http.ResponseWriter, content string) {
	fmt.Fprintf(w, content)
}

func JsonEncode(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
