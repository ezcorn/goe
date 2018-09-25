package goe

//
var serverManage = make(map[string][]func())

//
const manageStatus = "status"

//
const manageListen = "listen"

//
const manageAction = "action"

//
const manageRelate = "relate"

func initServerManage() {
	queue := []string{manageStatus, manageListen, manageAction, manageRelate}
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(serverManage[queue[i]]); j++ {
			serverManage[queue[i]][j]()
		}
	}
}

//
func joinManage(t string, f func()) {
	serverManage[t] = append(serverManage[t], f)
}
