package goe

import (
	"os"
	"time"
)

const vendor = "vendor"

type Input interface{}

type Output interface{}

func initDependTask(repository string) {
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
	gitClone(repository, vendor)
}

func CallMethod(depend string, action string, input Input) Output {
	return nil
}
