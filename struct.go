package goe

import (
	"net/http"
)

//
const runtimeStatus = "status"

//
const runtimeAction = "action"

//
const runtimeListen = "listen"

//
const runtimeRelate = "relate"

//
var registerRuntime = make(map[string][]func())

//
var serverConfig Config

//
var statusRegistry = make(map[int]func() string)

//
var listenRegistry = make(map[string]*ListenRegister)

//
var actionRegistry = make(Actions)

//
type Execute = func(http.ResponseWriter, *http.Request)

//
type ExecRet = func(http.ResponseWriter, *http.Request) Execute

//
type Actions map[string]*Action

//
type Listens []*Listen

type Config struct {
	Name    string  `json:"name"`    //
	Author  string  `json:"author"`  //
	Gateway string  `json:"gateway"` //
	Vendor  string  `json:"vendor"`  //
	Actions Actions `json:"actions"` //
}

type Action struct {
	Route   string  `json:"-"`       //
	Method  string  `json:"method"`  //
	Comment string  `json:"comment"` //
	Listens Listens `json:"listens"` //
	Execute Execute `json:"-"`       //
}

type Listen struct {
	Name    string  `json:"name"`    //
	Comment string  `json:"comment"` //
	Execute ExecRet `json:"-"`       //
}

type ListenRegister struct {
	Name    string  `json:"-"`       //
	Comment string  `json:"comment"` //
	Actions Actions `json:"actions"` //
	Listen  *Listen `json:"-"`       //
}
