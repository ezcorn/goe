package goe

import (
	"net/http"
)

//
type Execute = func(http.ResponseWriter, *http.Request)

//
type ExecRet = func(http.ResponseWriter, *http.Request) Execute

//
type Actions map[string]Action

//
type Listens map[string]Listen

type Config struct {
	Name    string  `json:"name"`    //
	Author  string  `json:"author"`  //
	Gateway string  `json:"gateway"` //
	Vendor  string  `json:"vendor"`  //
	Actions Actions `json:"actions"` //
}

type Action struct {
	Name    string  `json:"name"`    //
	Method  string  `json:"method"`  //
	Comment string  `json:"comment"` //
	Listens Listens `json:"listens"` //
	Execute Execute `json:"-"`       //
}

type Listen struct {
	Name    string  `json:"name"`    //
	Comment string  `json:"comment"` //
	Actions Actions `json:"actions"` //
	Execute ExecRet `json:"-"`       //
}

func Include(listen string, action string) {
	if listenObj, exist := ListenRegistry[listen]; exist {
		if actionObj, exist := ActionRegistry[action]; exist {
			ListenRegistry[listen].Actions[action] = actionObj
			ActionRegistry[action].Listens[listen] = listenObj
		}
	}
}
