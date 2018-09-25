package goe

import (
	"io/ioutil"
	"log"
)

//
var serverConfig Config

type Config struct {
	Name    string  `json:"name"`    //
	Author  string  `json:"author"`  //
	Gateway string  `json:"gateway"` //
	Vendor  string  `json:"vendor"`  //
	Actions Actions `json:"actions"` //
}

func initServerConfig() {

}

func readFile(fn string) []byte {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Panicln("File not found " + fn)
	}
	return buf
}
