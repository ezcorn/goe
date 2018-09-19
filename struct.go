package goe

import (
	"net/http"
)

//
type Execute = func(http.ResponseWriter, *http.Request)

//
type ExecRet = func(http.ResponseWriter, *http.Request) Execute

//
type Actions map[string]*Action

//
type Listens []*Listen

type Config struct {
	Name    string  `json:"name"`    //
	Author  string  `json:"author"`  //
	Gateway string  `json:"gateway"` //
	Vendor  string  `json:"vendor"`  //
	Actions Actions `json:"actions"` //
}

type Action struct {
	Route   string  `json:"-"`       //
	Comment string  `json:"comment"` //
	Method  string  `json:"method"`  //
	Listens Listens `json:"listens"` //
	Execute Execute `json:"-"`       //
}

type Listen struct {
	Name    string  `json:"name"`    //
	Comment string  `json:"comment"` //
	Execute ExecRet `json:"-"`       //
}

type ListenRegister struct {
	Name    string  `json:"-"`       //
	Comment string  `json:"comment"` //
	Actions Actions `json:"actions"` //
	Listen  *Listen `json:"-"`       //
}

func NewAction(route string, comment string, method string, execute Execute) *Action {
	if execute == nil {
		execute = func(writer http.ResponseWriter, request *http.Request) {}
	}
	return &Action{
		Route:   route,
		Method:  method,
		Comment: comment,
		Listens: Listens{},
		Execute: execute,
	}
}

func RegAction(new func() *Action, last func(a *Action)) {
	action := new()
	ActionRegistry[action.Route] = action
	last(action)
}

func NewListen(name string, comment string, execute ExecRet) *Listen {
	if execute == nil {
		execute = func(writer http.ResponseWriter, request *http.Request) Execute {
			return nil
		}
	}
	return &Listen{
		Name:    name,
		Comment: comment,
		Execute: execute,
	}
}

func RegListen(new func() *Listen, last func(l *Listen)) {
	listen := new()
	ListenRegistry[listen.Name] = &ListenRegister{
		Name:    listen.Name,
		Comment: listen.Comment,
		Actions: make(Actions),
		Listen:  listen,
	}
	last(listen)
}

func (listen *Listen) Append(actionRoute string) {
	listenRegister, exist := ListenRegistry[listen.Name]
	if !exist {
		return
	}
	if action, exist := ActionRegistry[actionRoute]; exist {
		if _, exist := listenRegister.Actions[actionRoute]; exist {
			return
		}
		listenRegister.Actions[actionRoute] = action
		action.Listens = append(action.Listens, listen)
	}
}

func (action *Action) Listen(listenName string) {
	for i := 0; i < len(action.Listens); i++ {
		if action.Listens[i].Name == listenName {
			return
		}
	}
	if listenRegister, exist := ListenRegistry[listenName]; exist {
		listenRegister.Actions[action.Route] = action
		action.Listens = append(action.Listens, listenRegister.Listen)
	}
}
