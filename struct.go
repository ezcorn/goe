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

type Action struct {
	Method  string `json:"method"`
	Comment string `json:"comment"`
	// This action use listens
	Listens *list.List `json:"listens"`
	// Do something
	Execute func(http.ResponseWriter, *http.Request) `json:"-"`
}

type Listen struct {
	Comment string `json:"comment"`
	// Used this listen's actions
	Actions *list.List `json:"actions"`
	// Result is a http function or nil
	Execute func(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request) `json:"-"`
}

func (listen Listen) Include(action Action) {
	// Link listen to action
	action.Listens.PushBack(listen)
	// Link action to listen
	listen.Actions.PushBack(action)
}
