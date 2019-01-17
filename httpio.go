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
)

// 获取头文件BODY的BYTE编码
func (in In) Body() []byte {
	body, err := ioutil.ReadAll(in.r.Body)
	defer in.r.Body.Close()
	if err != nil {
		return nil
	}
	return body
}

// 获取头文件BODY的字符串
func (in In) BodyStr() string {
	body := in.Body()
	if body == nil {
		return ""
	}
	return string(body)
}

// 获取头文件对象
func (in In) BodyObj(v interface{}) {
	libs.Json.Decode(in.Body(), v)
}

// 获取头文件字典
func (in In) BodyMap() Map {
	mp := Map{}
	in.BodyObj(&mp)
	return mp
}

// 判定头文件字典键存在,然后执行
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

// 获取头文件数组
func (in In) BodyArr() Arr {
	var arr Arr
	in.BodyObj(&arr)
	return arr
}

// 向网页输出一段内容
func (out Out) Echo(v interface{}) {
	switch v.(type) {
	case string:
		out.w.Header().Set("Content-Type", "application/text")
		_, _ = fmt.Fprintf(out.w, v.(string))
		break
	case View:
		break
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
				output = libs.Json.Encode(outputMap)
				break
			default:
				output = libs.Json.Encode(v)
			}
			_, _ = fmt.Fprintf(out.w, output)
			break
		}
	}
}

// 设置当前页面状态码
func (out Out) Status(code int) {
	if f, ok := statusRegistry[code]; ok {
		if f != nil {
			http.Error(out.w, f(code), code)
			return
		}
	}
	out.Status(http.StatusNotFound)
}
