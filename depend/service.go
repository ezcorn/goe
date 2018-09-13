package depend

type service interface {
	init()
}

func serviceParser() {

}

func Initialize() {
	configParser()
	serviceParser()
}
