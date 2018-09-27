package goe

const (
	manageStatus = "status"
	manageListen = "listen"
	manageAction = "action"
	manageRelate = "relate"
)

var (
	serverManage = make(map[string][]func())
)

func initServerManage() {
	queue := []string{manageStatus, manageListen, manageAction, manageRelate}
	for i := 0; i < len(queue); i++ {
		for j := 0; j < len(serverManage[queue[i]]); j++ {
			serverManage[queue[i]][j]()
		}
	}
}

func joinManage(t string, f func()) {
	serverManage[t] = append(serverManage[t], f)
}
