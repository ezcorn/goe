package goe

var ListenRegistry = make(map[string]*ListenRegister)

var ActionRegistry = make(Actions)

var GlobalConfig Config
