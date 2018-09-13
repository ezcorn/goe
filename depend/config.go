package depend

type config interface {
	parse() bool
}

type serviceRepoConfig struct{}

func (config serviceRepoConfig) parse() bool {
	// git clone service repo
	// check service repo
	return true
}

func configParser() {

}
