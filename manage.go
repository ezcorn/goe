package goe

//
const runtimeStatus = "status"

//
const runtimeAction = "action"

//
const runtimeListen = "listen"

//
const runtimeRelate = "relate"

func initServerManage() {
	queue := []string{runtimeStatus, runtimeListen, runtimeAction, runtimeRelate}
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(registerRuntime[queue[i]]); j++ {
			registerRuntime[queue[i]][j]()
		}
	}
}
