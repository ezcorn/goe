package goe

import (
	"os"
	"time"
)

func InitializeService(serviceRepo string) {
	go func() {
		duration := time.Second * 10
		for {
			// Clone serviceRepo to local
			os.RemoveAll("service")
			GitClone(serviceRepo, "service")
			// Read service json to memory

			// Sleep some time
			time.Sleep(duration)
		}
	}()
}
