package goe

import (
	"net/http"
	"strconv"
)

type RouteMap map[string]func(http.ResponseWriter, *http.Request)

func InitServer(routeMap RouteMap, port int, repository string) {
	initDependTask(repository)
	http.HandleFunc("/goe/makeDependFile", makeDependFile)
	for route, action := range routeMap {
		http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			action(writer, request)
		})
	}
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func makeDependFile(writer http.ResponseWriter, request *http.Request) {

}
