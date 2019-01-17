package goe

import "log"

// 输出分组日志
func LogPrintln(group string, info string) {
	log.Println("[ " + group + " ] : " + info)
}

//func initSystemApis() {
//	RegAction(func() *Action {
//		return NewAction("/apis/supplier", "", []string{
//			http.MethodPost,
//		}, func(in In, out Out) {
//			out.Echo(Norm{Data: map[string]interface{}{
//				"name":    currentServer.Name,
//				"host":    currentServer.Host,
//				"port":    strconv.Itoa(currentServer.Port),
//				"servers": currentServer.Servers,
//				"vendors": currentServer.Vendors,
//			}})
//		})
//	})
//	RegAction(func() *Action {
//		return NewAction("/apis/register", "", []string{
//			http.MethodPost,
//		}, func(in In, out Out) {
//
//		})
//	})
//}
