package goe

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func InitServer(port int) {
	// Read goe.json file
	json.Unmarshal(ReadFile("config.json"), &GlobalConfig)
	// Init vendor task
	initVendorTask(GlobalConfig.Vendor)
	for route, action := range GlobalAction {
		// Register restful api
		http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			// Execute listen list
			for _, listen := range action.Listens {
				// If listen is not true
				result := listen.Execute(writer, request)
				if result != nil {
					// If result is a function, Listen break
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
