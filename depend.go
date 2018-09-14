package goe

func InitializeDepend(repo string) {
	GitClone(repo, "depend")
}
