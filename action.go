package goe

import "net/http"

//
var actionRegistry = make(Actions)

//
type Actions map[string]*Action

//
type Program = func(w http.ResponseWriter, r *http.Request)

//
type Methods = []string

type Action struct {
	Route   string  `json:"-"`       //
	Method  Methods `json:"method"`  //
	Comment string  `json:"comment"` //
	Listens Listens `json:"listens"` //
	Program Program `json:"-"`       //
}

func (action *Action) Relate(listenName string) {
	for i := 0; i < len(action.Listens); i++ {
		if action.Listens[i].Name == listenName {
			return
		}
	}
	if listenRegister, exist := listenRegistry[listenName]; exist {
		listenRegister.Actions[action.Route] = action
		action.Listens = append(action.Listens, listenRegister.Listen)
	}
}

func NewAction(route string, comment string, method []string, program Program) *Action {
	if program == nil {
		program = func(w http.ResponseWriter, r *http.Request) {}
	}
	return &Action{
		Route:   route,
		Method:  method,
		Comment: comment,
		Listens: Listens{},
		Program: program,
	}
}

func RegAction(new func() *Action, relate func(a *Action)) {
	action := new()
	if new != nil {
		joinManage(manageAction, func() {
			actionRegistry[action.Route] = action
		})
		if relate != nil {
			joinManage(manageRelate, func() { relate(action) })
		}
	}
}
