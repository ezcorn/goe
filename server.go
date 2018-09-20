package goe

import (
	"log"
	"net/http"
	"strconv"
)

func InitServer(port int) {
	// TODO: Read config file
	// json.Unmarshal(readFile("config.json"), &serverConfig)
	// initVendorTask(serverConfig.Vendor)

	// TODO: Exec runtime function
	queue := []string{runtimeReg, runtimeAdd}
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(registerRuntime[runtimeReg]); j++ {
			registerRuntime[queue[i]][j]()
		}
	}

	// TODO: Register to handle
	for route, action := range actionRegistry {
		log.Println(`Action["` + route + `"]` + "\t" + jsonEncode(action))
		http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			for _, listen := range action.Listens {
				result := listen.Execute(writer, request)
				if result != nil {
					result(writer, request)
					break
				}
			}
			action.Execute(writer, request)
		})
	}

	// TODO: Register goe apis
	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {

	})

	// TODO: Start server
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
