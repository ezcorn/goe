package goe

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GitClone(repo string, name string) {
	_, err := exec.Command("git", "clone", repo, name).Output()
	cmd := "git clone " + repo + " " + name
	if err != nil {
		log.Panicln(cmd)
		os.Exit(1)
	} else {
		fmt.Println(cmd)
	}
}
