package goe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	In struct {
		r *http.Request
	}
	Out struct {
		w http.ResponseWriter
	}
)

func (in In) Body() interface{} {
	body, err := ioutil.ReadAll(in.r.Body)
	defer in.r.Body.Close()
	if err != nil {
		body = []byte{}
	}
	var result interface{}
	err = json.Unmarshal(body, &result)
	if result == nil {
		return ""
	} else {
		return result
	}
}

func (out Out) Echo(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Fprintf(out.w, v.(string))
		break
	default:
		out.w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(v)
		fmt.Fprintf(out.w, string(j))
	}
}

func (out Out) Status(code int) {
	if f, ok := statusRegistry[code]; ok {
		if f != nil {
			http.Error(out.w, f(code), code)
		}
	}
	out.Status(http.StatusNotFound)
}
