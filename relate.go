package goe

func RelateActionToListen(actionRoute string, listenName string) {
	if action, actionExist := actionRegistry[actionRoute]; actionExist {
		if listenRegister, listenExist := listenRegistry[listenName]; listenExist {
			joinManage(manageRelate, func() {
				listenRegister.Actions[action.Route] = action
				action.Listens = append(action.Listens, listenRegister.Listen)
			})
		}
	}
}
