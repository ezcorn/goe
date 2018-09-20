package goe

import (
	"time"
)

const vendor = "vendor"

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
