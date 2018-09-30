package goe

type (
	Actions map[string]*Action
	Program = func(in In, out Out)
	Methods = []string
	Action  struct {
		Route   string  `json:"-"`       //
		Method  Methods `json:"method"`  //
		Comment string  `json:"comment"` //
		Listens Listens `json:"listens"` //
		Program Program `json:"-"`       //
	}
)

var (
	actionRegistry = make(Actions)
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
		program = func(in In, out Out) {}
	}
	return &Action{
		Route:   route,
		Method:  method,
		Comment: comment,
		Listens: Listens{},
		Program: program,
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
