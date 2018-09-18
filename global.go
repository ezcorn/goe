package goe

import "container/list"

var GlobalServerListen = list.New()

var GlobalServerRouter = make(Router)

var GlobalServerConfig Config
