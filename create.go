package goe

import "net/http"

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

func RegAction(new func() *Action, relate func(a *Action)) {
	action := new()
	if new != nil {
		joinRuntime(runtimeAction, func() {
			actionRegistry[action.Route] = action
		})
		if relate != nil {
			joinRuntime(runtimeRelate, func() { relate(action) })
		}
	}
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

func RegListen(new func() *Listen, relate func(l *Listen)) {
	listen := new()
	if new != nil {
		joinRuntime(runtimeListen, func() {
			listenRegistry[listen.Name] = &ListenRegister{
				Name:    listen.Name,
				Comment: listen.Comment,
				Actions: make(Actions),
				Listen:  listen,
			}
		})
		if relate != nil {
			joinRuntime(runtimeRelate, func() { relate(listen) })
		}
	}
}

func RegStatus(code int, f func() string) {
	if f != nil {
		joinRuntime(runtimeStatus, func() {
			statusRegistry[code] = f
		})
	}
}
