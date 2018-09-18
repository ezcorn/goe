package goe

import (
	"log"
	"os"
	"os/exec"
	"time"
)

const vendor = "vendor"

type Input interface{}

type Output interface{}

func initVendorTask(repository string) {
	cloneVendor(repository)
	go func() {
		duration := time.Minute * 5
		for {
			time.Sleep(duration)
			cloneVendor(repository)
		}
	}()
}

func cloneVendor(repository string) {
	os.RemoveAll(vendor)
	_, err := exec.Command("git", "clone", repository, vendor).Output()
	log.Println("git clone " + repository + " " + vendor)
	if err != nil {
		os.Exit(1)
	}
}

func CallVendor(name string, action string, input Input) Output {
	return nil
}
