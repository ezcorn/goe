package libs

import (
	"io/ioutil"
	"os"
)

const (
	filePermission = 0755
)

type (
	IO struct{}
)

func (IO) Exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func (io IO) MkDir(dir string) {
	if !io.Exist(dir) {
		os.Mkdir(dir, filePermission)
	}
}

func (IO) Write(fileName string, content string) {
	ioutil.WriteFile(fileName, []byte(content), filePermission)
}
