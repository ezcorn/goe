package goe

import (
	"time"
)

func initVendorTask() {
	go func() {
		duration := time.Minute * 5
		for {
			time.Sleep(duration)
		}
	}()
}

func CallVendor(vendorName string, actionRoute string, input interface{}) interface{} {
	return nil
}
