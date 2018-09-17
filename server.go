package goe

import (
	"log"
	"net/http"
	"strconv"
)

type RouteMap map[string]func(http.ResponseWriter, *http.Request)

var HttpRouteMap RouteMap = make(RouteMap)

func InitServer(port int, repository string) {
	initVendorTask(repository)
	http.HandleFunc("/goe/makeDependFile", makeDependFile)
	for route, action := range HttpRouteMap {
		log.Println("server route " + route)
		http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			action(writer, request)
		})
	}
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func makeDependFile(writer http.ResponseWriter, request *http.Request) {

}
