package goe

type (
	Map     map[string]interface{}
	Arr     []interface{}
	Program = func(in In, out Out, libs Libs)
	Process = func(in In) Program
	Methods = []string

	Action struct {
		Route   string  `json:"-"`       //
		Method  Methods `json:"method"`  //
		Comment string  `json:"comment"` //
		Listens Listens `json:"listens"` //
		Program Program `json:"-"`       //
	}
	Actions map[string]*Action

	Listen struct {
		Name    string  `json:"name"`    //
		Comment string  `json:"comment"` //
		Process Process `json:"-"`       //
	}
	Listens []*Listen

	ListenRegister struct {
		Name    string  `json:"-"`       //
		Comment string  `json:"comment"` //
		Actions Actions `json:"actions"` //
		Listen  *Listen `json:"-"`       //
	}
)

var (
	actionRegistry = make(Actions)
	listenRegistry = make(map[string]*ListenRegister)
)

func (action *Action) MethodContains(method string) bool {
	for _, v := range action.Method {
		if v == method {
			return true
		}
	}
	return false
}

func NewAction(route string, comment string, method []string, program Program) *Action {
	if program == nil {
		program = func(in In, out Out, libs Libs) {}
	}
	return &Action{
		Route:   route,
		Method:  method,
		Comment: comment,
		Listens: Listens{},
		Program: program,
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

func RegAction(new func() *Action) {
	action := new()
	if new != nil {
		joinManage(manageAction, func() {
			actionRegistry[action.Route] = action
		})
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

func RelateActionToListen(actionRoute string, listenName string) {
	joinManage(manageRelate, func() {
		if action, actionExist := actionRegistry[actionRoute]; actionExist {
			if listenRegister, listenExist := listenRegistry[listenName]; listenExist {
				listenRegister.Actions[action.Route] = action
				action.Listens = append(action.Listens, listenRegister.Listen)
			}
		}
	})
}
