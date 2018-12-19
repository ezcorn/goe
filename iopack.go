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

var (
	IO io
)

func (io) Exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func (io io) Dirs(dirPath string) []string {
	if fileInfoList, err := ioutil.ReadDir(dirPath); err == nil {
		var result []string
		for _, fileInfo := range fileInfoList {
			if fileInfo.IsDir() {
				result = append(result, fileInfo.Name())
			}
		}
		return result
	}
	return nil
}

func (io io) MkDir(dir string) {
	if !io.Exist(dir) {
		_ = os.Mkdir(dir, filePermission)
	}
}

func (io) Write(fileName string, content string) {
	_ = ioutil.WriteFile(fileName, []byte(content), filePermission)
}

func (io io) WriteJson(fileName string, data interface{}) {
	io.Write(fileName, jsonEncode(data))
}

func (io io) Read(fileName string) []byte {
	if io.Exist(fileName) {
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			goto RET
		}
		return b
	}
RET:
	return nil
}

func (io io) ReadJson(fileName string, v interface{}) {
	str := io.Read(fileName)
	if str != nil {
		jsonDecode(str, v)
	}
}
