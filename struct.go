package goe

import (
	"container/list"
	"net/http"
)

type Router map[string]Action

type Config struct {
	Name    string `json:"name"`
	Author  string `json:"author"`
	Gateway string `json:"gateway"`
	Vendor  string `json:"vendor"`
	Router  Router `json:"router"`
}

/**

 */
type Action struct {
	Listens *list.List
	Execute func(http.ResponseWriter, *http.Request)
}

/**

 */
type Listen struct {
	Actions *list.List
	Execute func(http.ResponseWriter, *http.Request) bool
}

func (listen Listen) include(action Action) {
	action.Listens.PushBack(listen)
	listen.Actions.PushBack(action)
}
