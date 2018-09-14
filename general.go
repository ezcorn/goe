package goe

import (
	"log"
	"os"
	"os/exec"
)

func GitClone(repo string, name string) {
	_, err := exec.Command("git", "clone", repo, name).Output()
	if err != nil {
		log.Panicln("git clone " + repo + " " + name)
		os.Exit(1)
	}
}

func GitPull(repo string) {

}
