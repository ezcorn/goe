package goe

import (
	"os"
	"time"
)

func initDependTask(repository string) {
	go func() {
		duration := time.Minute * 5
		for {
			// Clone serviceRepo to local
			os.RemoveAll("depend")
			gitClone(repository, "depend")
			// Read service json to memory

			// Sleep some time
			time.Sleep(duration)
		}
	}()
}
