package goe

type (
	root struct{}
)

var (
	// getRecord  []string
	groupCache = make(map[string][]interface{})
)

func (root) Set(group string, data interface{}) {
	// MEM -> LOCAL STO -> PRIME STO -> CENTER STO

	// HTML5 -> ACCESS -> SERVER
	// HTML5 <- ACCESS <- SERVER
	//      DATA      DATA
}

func (root) Get(group string, where map[string]string) {
	// MEM <- MEM | MEM <- STO
	if _, ok := groupCache[group]; ok {

	} else {

	}
}
