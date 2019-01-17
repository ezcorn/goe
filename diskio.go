package goe

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	filePermission = 0755
	pathCharacter  = "/"
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

// 层叠创建文件夹
func (io io) MkDir(dir string) {
	ds := strings.Split(dir, pathCharacter)
	length := len(ds)
	if length > 0 {
		buf := ds[0]
		for k, d := range ds {
			d = strings.TrimSpace(d)
			if d != "" {
				if k > 0 {
					buf += pathCharacter + d
				}
				if !io.Exist(buf) {
					err := os.Mkdir(buf, filePermission)
					if err != nil {
						log.Println(err.Error())
						return
					}
				}
			}
		}
	}
}

// 创建一个文件
func (io io) Create(fileName string) {
	if !io.Exist(fileName) {
		fs := strings.Split(fileName, pathCharacter)
		length := len(fs)
		if length > 0 {
			if length > 1 {
				io.MkDir(strings.Join(fs[:length-1], pathCharacter))
			}
			f, _ := os.Create(fileName)
			defer f.Close()
		}
	}
}

func (io) Write(fileName string, content string) {
	_ = ioutil.WriteFile(fileName, []byte(content), filePermission)
}

func (io) WriteAppend(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, filePermission)
	if err == nil {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

func (io io) WriteAppendLine(fileName string, content string) error {

	return io.WriteAppend(fileName, content+"\n")
}

func (io io) WriteJson(fileName string, data interface{}) string {
	json := JSON.EncodeIndent(data)
	io.Write(fileName, json)
	return json
}

//func (io io) WriteJsonBase64(fileName string, data interface{}) string {
//
//}

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
		JSON.Decode(str, v)
	}
}
