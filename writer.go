package goe

import (
	"fmt"
	"net/http"
)

func Echo(w http.ResponseWriter, content string) {
	fmt.Fprintf(w, content)
}
