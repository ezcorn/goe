package goe

import (
	"net/http"
)

//
var registerRuntime = make(map[string][]func())

//
var serverConfig Config

//
var statusRegistry = make(map[int]func(int) string)

//
var listenRegistry = make(map[string]*ListenRegister)

//
var actionRegistry = make(Actions)

//
type Program = func(w http.ResponseWriter, r *http.Request)

//
type Process = func(w http.ResponseWriter, r *http.Request) Program

//
type Actions map[string]*Action

//
type Listens []*Listen

//
type Methods = []string

type Config struct {
	Name    string  `json:"name"`    //
	Author  string  `json:"author"`  //
	Gateway string  `json:"gateway"` //
	Vendor  string  `json:"vendor"`  //
	Actions Actions `json:"actions"` //
}

type Action struct {
	Route   string  `json:"-"`       //
	Method  Methods `json:"method"`  //
	Comment string  `json:"comment"` //
	Listens Listens `json:"listens"` //
	Program Program `json:"-"`       //
}

type Listen struct {
	Name    string  `json:"name"`    //
	Comment string  `json:"comment"` //
	Process Process `json:"-"`       //
}

type ListenRegister struct {
	Name    string  `json:"-"`       //
	Comment string  `json:"comment"` //
	Actions Actions `json:"actions"` //
	Listen  *Listen `json:"-"`       //
}
