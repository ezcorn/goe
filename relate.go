package goe

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
