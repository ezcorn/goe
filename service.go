package goe

func InitializeService(serviceRepo string) {
	// Clone serviceRepo to local
	GitClone(serviceRepo, "service")
	// Start serviceRepo refresh schedule

}
