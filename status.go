package goe

import (
	"net/http"
	"strconv"
)

var statusRegistry = make(map[int]func(int) string)

var httpStatus = []int{
	100, 101, 102,
	200, 201, 202, 203, 204, 205, 206, 207, 208, 226,
	300, 301, 302, 303, 304, 305, 307, 308,
	400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 426, 428, 429, 431, 451,
	500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511,
}

func initServerStatus() {
	for i := 0; i < len(httpStatus); i++ {
		RegStatus(httpStatus[i], func(code int) string {
			return strconv.Itoa(code) + " " + http.StatusText(code)
		})
	}
}

func httpState(w http.ResponseWriter, code int) {
	if f, ok := statusRegistry[code]; ok {
		if f != nil {
			http.Error(w, f(code), code)
			return
		}
	}
	httpState(w, http.StatusNotFound)
}

func RegStatus(code int, f func(int) string) {
	if f != nil {
		joinManage(manageStatus, func() {
			statusRegistry[code] = f
		})
	}
}
