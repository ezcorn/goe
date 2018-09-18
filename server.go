package goe

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func InitServer(port int) {
	// Read goe.json file
	json.Unmarshal(ReadFile("config.json"), &GlobalServerConfig)
	// Init vendor task
	initVendorTask(GlobalServerConfig.Vendor)
	for route, action := range GlobalServerRouter {
		// Register restful api
		http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			// Execute listen list
			for e := action.Listens.Front(); e != nil; e = e.Next() {
				// If listen is not true
				result := e.Value.(*Listen).Execute(writer, request)
				// If result is a function, Listen break
				if result != nil {
					result(writer, request)
					break
				}
			}
			// Execute action
			action.Execute(writer, request)
		})
	}
	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {

	})
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
