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
type Listens []Listen

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

func NewAction(name string, comment string, method string, execute Execute) Action {
	if execute == nil {
		execute = func(writer http.ResponseWriter, request *http.Request) {}
	}
	return Action{
		Name:    name,
		Method:  method,
		Comment: comment,
		Listens: Listens{},
		Execute: execute,
	}
}

type Listen struct {
	Name    string  `json:"name"`    //
	Comment string  `json:"comment"` //
	Actions Actions `json:"-"`       //
	Execute ExecRet `json:"-"`       //
}

func NewListen(name string, comment string, execute ExecRet) Listen {
	if execute == nil {
		execute = func(writer http.ResponseWriter, request *http.Request) Execute {
			return nil
		}
	}
	return Listen{
		Name:    name,
		Comment: comment,
		Actions: make(Actions),
		Execute: execute,
	}
}

func (listen Listen) Join(action string) {
	//if listenObj, exist := ListenRegistry[listen]; exist {
	//	if actionObj, exist := ActionRegistry[action]; exist {
	//		ListenRegistry[listen].Actions[action] = actionObj
	//		ActionRegistry[action].Listens[listen] = listenObj
	//	}
	//}
}
