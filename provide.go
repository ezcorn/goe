package goe

import (
	"net/http"
	"strconv"
)

type HttpRoute map[string]func(http.ResponseWriter, *http.Request)

func InitProvideServer(routes HttpRoute, port int) {
	for route, action := range routes {
		http.HandleFunc(route, action)
	}
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
