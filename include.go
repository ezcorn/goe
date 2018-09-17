package goe

import (
	"os"
	"time"
)

func initIncludeServer(includeRepo string) {
	go func() {
		duration := time.Minute * 5
		for {
			// Clone serviceRepo to local
			os.RemoveAll("include")
			GitClone(includeRepo, "include")
			// Read service json to memory

			// Sleep some time
			time.Sleep(duration)
		}
	}()
}
