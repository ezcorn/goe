package goe

type (
	Listens []*Listen
	Process = func(in In) Program
	Listen  struct {
		Name    string  `json:"name"`    //
		Comment string  `json:"comment"` //
		Process Process `json:"-"`       //
	}
	ListenRegister struct {
		Name    string  `json:"-"`       //
		Comment string  `json:"comment"` //
		Actions Actions `json:"actions"` //
		Listen  *Listen `json:"-"`       //
	}
)

var (
	listenRegistry = make(map[string]*ListenRegister)
)

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

func RegListen(new func() *Listen) {
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
	}
}
