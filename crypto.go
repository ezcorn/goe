package goe

import (
	"crypto/md5"
	"fmt"
)

type (
	crypto struct{}
)

var (
	Crypto crypto
)

func (crypto) MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

//func (encrypt) MD58(str string) string {
//
//}
