package goe

var ListenRegistry = make(map[Listen]map[string]Action)

var ActionRegistry = make(Actions)

var GlobalConfig Config
