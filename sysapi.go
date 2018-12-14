package goe

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
