package goe

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func readFile(fn string) []byte {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Panicln("File not found " + fn)
	}
	return buf
}

//
func gitClone(repo string, name string) {
	os.RemoveAll(name)
	_, err := exec.Command("git", "clone", repo, name).Output()
	log.Println("git clone " + repo + " " + name)
	if err != nil {
		os.Exit(1)
	}
}

//
func jsonEncode(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

//
func joinRuntime(t string, f func()) {
	registerRuntime[t] = append(registerRuntime[t], f)
}
