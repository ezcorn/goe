package goe

type (
	// 控制器函数模型
	program = func(in In, out Out)
	// 监听器函数模型
	process = func(in In) program
	// 可访问的方法数组
	methods = []string
	// 控制器结构
	Action struct {
		Route   string  `json:"-"`       //
		Method  methods `json:"method"`  //
		Comment string  `json:"comment"` //
		Listens listens `json:"listens"` //
		Program program `json:"-"`       //
	}
	// 控制器数组
	actions map[string]*Action
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

// 该ACTION中是否存在请求模式
func (action *Action) methodContains(method string) bool {
	for _, v := range action.Method {
		if v == method {
			return true
		}
	}
	return false
}

// 加入协调管理
func joinManage(t string, f func()) {
	serverManage[t] = append(serverManage[t], f)
}

func (s Server) RegStatus(code int, f func(int) string) {
	if f != nil {
		joinManage(manageStatus, func() { statusRegistry[code] = f })
	}
}

func (Server) RegAction(new func() *Action) {
	action := new()
	if new != nil {
		joinManage(manageAction, func() {
			actionRegistry[action.Route] = action
		})
	}
}

func (Server) RegListen(new func() *listen) {
	listen := new()
	if new != nil {
		joinManage(manageListen, func() {
			listenRegistry[listen.Name] = &listenRegister{
				Name:    listen.Name,
				Comment: listen.Comment,
				Actions: make(actions),
				Listen:  listen,
			}
		})
	}
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
