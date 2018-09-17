package goe

import (
	"net/http"
	"strconv"
)

func InitProvideServer(routes map[string]func(http.ResponseWriter, *http.Request), port int) {
	for route, action := range routes {
		http.HandleFunc(route, action)
	}
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func MakeEmptyRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return make(map[string]func(http.ResponseWriter, *http.Request))
}
