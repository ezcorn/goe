package goe

import (
	"net/http"
)

type NormFunc = func(http.ResponseWriter, *http.Request)

type NormFuncListen = func(http.ResponseWriter, *http.Request) NormFunc

type Config struct {
	Name    string            `json:"name"`
	Author  string            `json:"author"`
	Gateway string            `json:"gateway"`
	Vendor  string            `json:"vendor"`
	Actions map[string]Action `json:"actions"`
}

type Action struct {
	Name    string            `json:"name"`
	Method  string            `json:"method"`
	Comment string            `json:"comment"`
	Listens map[string]Listen `json:"listens"`
	Execute NormFunc          `json:"-"`
}

type Listen struct {
	Name    string            `json:"name"`
	Comment string            `json:"comment"`
	Actions map[string]Action `json:"actions"`
	Execute NormFuncListen    `json:"-"`
}

func (listen Listen) Include(route string) {
	// Link listen and action
	if action, exist := GlobalAction[route]; exist {
		// Add action to listen
		listen.Actions[route] = action
		// Add listen to action
		action.Listens[listen.Name] = listen
	}
}
