package goe

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	jsonOutputCode200 = 200
	jsonOutputCode500 = 500
)

type (
	// HTTP输入封装
	In struct {
		r *http.Request
	}
	// HTTP输出封装
	Out struct {
		w http.ResponseWriter
	}
	// 标准JSON输入
	Norm struct {
		Data interface{}
		Info string
	}
	// 输出界面
	View struct {
	}
	// 工具库
	Libs struct {
		IO   io
		Root root
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

func (in In) BodyStr() string {
	body := in.Body()
	if body == nil {
		return ""
	}
	return string(body)
}

func (in In) BodyObj(v interface{}) {
	jsonDecode(in.Body(), v)
}

func (in In) BodyMap() Map {
	mp := Map{}
	in.BodyObj(&mp)
	return mp
}

func (in In) BodyMapKeyExist(keys []string, exec func(body Map)) {
	mp := in.BodyMap()
	if mp != nil {
		for _, k := range keys {
			if _, existKey := mp[k]; !existKey {
				return
			}
		}
	}
	exec(mp)
}

func (in In) BodyArr() Arr {
	var arr Arr
	in.BodyObj(&arr)
	return arr
}

func (out Out) Echo(v interface{}) {
	switch v.(type) {
	case string:
		out.w.Header().Set("Content-Type", "application/text")
		_, _ = fmt.Fprintf(out.w, v.(string))
		break
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
				norm := v.(Norm)
				outputMap := Map{
					"code": jsonOutputCode200,
					"data": norm.Data,
					"info": norm.Info,
				}
				if norm.Info != "" {
					outputMap["code"] = jsonOutputCode500
				}
				output = jsonEncode(outputMap)
				break
			default:
				output = jsonEncode(v)
			}
			_, _ = fmt.Fprintf(out.w, output)
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
