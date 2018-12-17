package goe

type (
	// 控制器函数模型
	program = func(in In, out Out, libs Libs)
	// 监听器函数模型
	process = func(in In, libs Libs) program
	// 可访问的方法数组
	methods = []string
	// 控制器结构
	action struct {
		Route   string  `json:"-"`       //
		Method  methods `json:"method"`  //
		Comment string  `json:"comment"` //
		Listens listens `json:"listens"` //
		Program program `json:"-"`       //
	}
	// 控制器数组
	actions map[string]*action
	// 监听器结构
	listen struct {
		Name    string  `json:"name"`    //
		Comment string  `json:"comment"` //
		Process process `json:"-"`       //
	}
	// 监听器数组
	listens []*listen
	// 监听器注册结构
	listenRegister struct {
		Name    string  `json:"-"`       //
		Comment string  `json:"comment"` //
		Actions actions `json:"actions"` //
		Listen  *listen `json:"-"`       //
	}
)

var (
	// 控制器注册表
	actionRegistry = make(actions)
	// 监听器注册表
	listenRegistry = make(map[string]*listenRegister)
)

//
func (action *action) methodContains(method string) bool {
	for _, v := range action.Method {
		if v == method {
			return true
		}
	}
	return false
}

func (Server) RegAction(route string, comment string, method []string, program program) {
	if program == nil {
		program = func(in In, out Out, libs Libs) {}
	}
	joinManage(manageAction, func() {
		actionRegistry[route] = &action{
			Route:   route,
			Method:  method,
			Comment: comment,
			Listens: listens{},
			Program: program,
		}
	})
}

func (Server) RegListen(name string, comment string, process process) {
	if process == nil {
		process = func(in In, libs Libs) program { return nil }
	}
	joinManage(manageListen, func() {
		listenRegistry[name] = &listenRegister{
			Name:    name,
			Comment: comment,
			Actions: make(actions),
			Listen: &listen{
				Name:    name,
				Comment: comment,
				Process: process,
			},
		}
	})
}

func (Server) RelateActionToListen(actionRoute string, listenName string) {
	joinManage(manageRelate, func() {
		if action, actionExist := actionRegistry[actionRoute]; actionExist {
			if listenRegister, listenExist := listenRegistry[listenName]; listenExist {
				listenRegister.Actions[action.Route] = action
				action.Listens = append(action.Listens, listenRegister.Listen)
			}
		}
	})
}
