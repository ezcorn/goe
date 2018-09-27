package goe

import (
	"encoding/json"
	"fmt"
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

func (out Out) Echo(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Fprintf(out.w, v.(string))
		break
	default:
		j, _ := json.Marshal(v)
		fmt.Fprintf(out.w, string(j))
	}
}

func (out Out) Status(code int) {
	if f, ok := statusRegistry[code]; ok {
		if f != nil {
			http.Error(out.w, f(code), code)
			return
		}
	}
	out.Status(http.StatusNotFound)
}
