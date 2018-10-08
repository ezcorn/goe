package goe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	filePermission = 0755
)

type (
	In struct {
		r *http.Request
	}
	Out struct {
		w http.ResponseWriter
	}
)

func (in In) Body() []byte {
	body, err := ioutil.ReadAll(in.r.Body)
	defer in.r.Body.Close()
	if err != nil {
		return nil
	}
	return body
}

func (out Out) Echo(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	switch v[0].(type) {
	case string:
		fmt.Fprintf(out.w, v[0].(string))
		break
	default:
		out.w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(map[string]interface{}{
			"code": "",
			"data": v[0],
			"info": "",
		})
		fmt.Fprintf(out.w, string(j))
	}
}

func (out Out) status(code int) {
	if f, ok := statusRegistry[code]; ok {
		if f != nil {
			http.Error(out.w, f(code), code)
			return
		}
	}
	out.status(http.StatusNotFound)
}

func (out Out) Status(b bool, code int) bool {
	if b {
		out.status(code)
	}
	return b
}

func (out Out) IoExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func (out Out) IoMkDir(dir string) {
	if !out.IoExist(dir) {
		os.Mkdir(dir, filePermission)
	}
}

func (out Out) IoWrite(fileName string, content string) {

}
