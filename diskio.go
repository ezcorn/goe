package goe

import (
	"io/ioutil"
	"os"
)

const (
	filePermission = 0755
)

type (
	io struct{}
)

func (io) Exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func (io io) MkDir(dir string) {
	if !io.Exist(dir) {
		_ = os.Mkdir(dir, filePermission)
	}
}

func (io) Write(fileName string, content string) {
	_ = ioutil.WriteFile(fileName, []byte(content), filePermission)
}
