package goe

import (
	"io/ioutil"
	"log"
)

func readFile(fn string) []byte {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Panicln("File not found " + fn)
	}
	return buf
}
