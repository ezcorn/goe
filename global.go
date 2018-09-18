package goe

var ListenRegistry = make(map[string]Listen)

var ActionRegistry = make(map[string]Action)

var GlobalConfig Config
