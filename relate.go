package goe

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
