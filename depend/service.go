package depend

type service interface {
	init()
}

func serviceParser() {

}

func DependInitialize() {
	configParser()
	serviceParser()
}
