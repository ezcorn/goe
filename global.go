package goe

var ListenRegistry = make(map[Listen]Actions)

var ActionRegistry = make(Actions)

var GlobalConfig Config
