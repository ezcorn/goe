package goe

import (
	"net/http"
	"strconv"
)

type RouteMap map[string]func(http.ResponseWriter, *http.Request)

func InitServer(routeMap RouteMap, port int, repository string) {
	initDependTask(repository)
	for route, action := range routeMap {
		http.HandleFunc(route, action)
	}
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
