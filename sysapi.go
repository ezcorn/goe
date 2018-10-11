package goe

import (
	"net/http"
	"strconv"
)

func initSystemApis() {
	RegAction(func() *Action {
		return NewAction("/apis/whoAmI", "create Action", []string{
			http.MethodPost,
		}, func(in In, out Out) {
			out.Echo(Norm{Data: map[string]interface{}{
				"name":    currentServer.Name,
				"host":    currentServer.Host,
				"port":    strconv.Itoa(currentServer.Port),
				"servers": currentServer.Servers,
				"vendors": currentServer.Vendors,
			}})
		})
	})
}
