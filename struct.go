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
	Route   string  `json:"route"`   //
	Method  string  `json:"method"`  //
	Comment string  `json:"comment"` //
	Listens Listens `json:"listens"` //
	Execute Execute `json:"-"`       //
}

type Listen struct {
	Name    string  `json:"name"`    //
	Comment string  `json:"comment"` //
	Execute ExecRet `json:"-"`       //
}

type ListenRegister struct {
	Listen  Listen  `json:"listen"`  //
	Actions Actions `json:"actions"` //
}

func NewAction(route string, comment string, method string, execute Execute) Action {
	if execute == nil {
		execute = func(writer http.ResponseWriter, request *http.Request) {}
	}
	return Action{
		Route:   route,
		Method:  method,
		Comment: comment,
		Listens: Listens{},
		Execute: execute,
	}
}

func AppendAction(action Action) {
	ActionRegistry[action.Route] = action
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
		Execute: execute,
	}
}

func (listen Listen) Join(action string) {
	listenRegister, exist := ListenRegistry[listen.Name]
	if !exist {
		listenRegister = ListenRegister{
			Listen:  listen,
			Actions: make(Actions),
		}
	}
	if actionObj, exist := ActionRegistry[action]; exist {
		listenRegister.Actions[action] = actionObj
		actionObj.Listens = append(actionObj.Listens, listen)
		ActionRegistry[action] = actionObj
	}
	ListenRegistry[listen.Name] = listenRegister
}

func (action Action) Listen(listen Listen) {

}
