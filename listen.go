package goe

//
var listenRegistry = make(map[string]*ListenRegister)

//
type Listens []*Listen

//
type Process = func(in In) Program

type Listen struct {
	Name    string  `json:"name"`    //
	Comment string  `json:"comment"` //
	Process Process `json:"-"`       //
}

type ListenRegister struct {
	Name    string  `json:"-"`       //
	Comment string  `json:"comment"` //
	Actions Actions `json:"actions"` //
	Listen  *Listen `json:"-"`       //
}

func (listen *Listen) Relate(actionRoute string) {
	listenRegister, exist := listenRegistry[listen.Name]
	if !exist {
		return
	}
	if action, exist := actionRegistry[actionRoute]; exist {
		if _, exist := listenRegister.Actions[actionRoute]; exist {
			return
		}
		listenRegister.Actions[actionRoute] = action
		action.Listens = append(action.Listens, listen)
	}
}

func NewListen(name string, comment string, process Process) *Listen {
	if process == nil {
		process = func(in In) Program { return nil }
	}
	return &Listen{
		Name:    name,
		Comment: comment,
		Process: process,
	}
}

func RegListen(new func() *Listen, relate func(l *Listen)) {
	listen := new()
	if new != nil {
		joinManage(manageListen, func() {
			listenRegistry[listen.Name] = &ListenRegister{
				Name:    listen.Name,
				Comment: listen.Comment,
				Actions: make(Actions),
				Listen:  listen,
			}
		})
		if relate != nil {
			joinManage(manageRelate, func() { relate(listen) })
		}
	}
}
