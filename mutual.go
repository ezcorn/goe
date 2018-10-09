package goe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	filePermission    = 0755
	jsonOutputCode200 = 200
	jsonOutputCode500 = 500
)

type (
	In struct {
		r *http.Request
	}
	Out struct {
		w http.ResponseWriter
	}
	Norm struct {
		Data interface{}
		Info string
	}
	View struct {
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

func (out Out) Echo(v interface{}) {
	switch v.(type) {
	case string:
		{
			fmt.Fprintf(out.w, v.(string))
			break
		}
	case View:
		{
			break
		}
	default:
		{
			out.w.Header().Set("Content-Type", "application/json")
			var output string
			switch v.(type) {
			case Norm:
				{
					norm := v.(Norm)
					code := jsonOutputCode200
					info := norm.Info
					data := norm.Data
					if info != "" {
						code = jsonOutputCode500
					}
					if data == nil {
						data = make(map[string]string)
					}
					output = out.JsonEncode(map[string]interface{}{
						"code": code,
						"data": data,
						"info": info,
					})
					break
				}
			default:
				{
					output = out.JsonEncode(v)
				}
			}
			fmt.Fprintf(out.w, output)
			break
		}
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
func (out Out) JsonEncode(data interface{}) string {
	j, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(j)
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
	ioutil.WriteFile(fileName, []byte(content), filePermission)
}
