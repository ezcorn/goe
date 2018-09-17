package goe

import (
	"log"
	"os"
	"os/exec"
)

func GitClone(repo string, name string) {
	_, err := exec.Command("git", "clone", repo, name).Output()
	cmd := "git clone " + repo + " " + name
	log.Println(cmd)
	if err != nil {
		os.Exit(1)
	}
}
