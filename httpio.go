package goe

import (
	"fmt"
	"github.com/ezcorn/goe/libs"
	"io/ioutil"
	"net/http"
)

const (
	jsonOutputCode200 = 200
	jsonOutputCode500 = 500
)

type (
	In struct {
		r    *http.Request
		libs Libs
	}
	Out struct {
		w    http.ResponseWriter
		libs Libs
	}
	Norm struct {
		Data interface{}
		Info string
	}
	View struct {
	}
	Libs struct {
		Queue libs.Queue
		Json  libs.Json
		IO    libs.IO
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
	in.libs.Json.Decode(in.Body(), v)
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
				output = out.libs.Json.Encode(outputMap)
				break
			default:
				output = out.libs.Json.Encode(v)
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
