package goe

import (
	"log"
	"os"
	"os/exec"
	"time"
)

const vendor = "vendor"

type Vendor struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Master Node   `json:"master"`
}

//type Input interface{}
//
//type Output interface{}

func initVendorTask(repository string) {
	gitClone(repository, vendor)
	go func() {
		duration := time.Minute * 5
		for {
			time.Sleep(duration)
			gitClone(repository, vendor)
		}
	}()
}

func CallVendor(vendorName string, actionRoute string, input interface{}) interface{} {
	return nil
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
