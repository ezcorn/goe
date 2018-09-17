package goe

import (
	"os"
	"time"
)

const local = "depend"

type Input interface{}

type Output interface{}

func initDependTask(repository string) {
	go func() {
		duration := time.Minute * 5
		for {
			// Clone serviceRepo to local
			os.RemoveAll(local)
			gitClone(repository, local)
			// Read service json to memory

			// Sleep some time
			time.Sleep(duration)
		}
	}()
}

func CallMethod(depend string, action string, input Input) Output {
	return nil
}
