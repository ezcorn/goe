package goe

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const name = "config.json"

//
var serverConfig Config

type Config struct {
	Detail  Vendor   `json:"detail"`
	Slaves  []Node   `json:"slaves"`
	Actions Actions  `json:"actions"`
	Vendor  []Vendor `json:"vendor"`
	DNS     string   `json:"dns"`
}

func initServerConfig() {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalln("Config not found")
	}
	serverConfig = Config{}
	err = json.Unmarshal(buf, &serverConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (config Config) addSlave(host string, port int) {

}

func (config Config) addAction(action Action) {

}

func (config Config) addVendor(name string, master string) {

}
