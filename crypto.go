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

// MD5加密
func (crypto) MD5(str string, size uint8) string {
	if size > 0 && size <= 32 {
		n := (32 - size) / 2
		data := []byte(str)
		has := md5.Sum(data)
		return fmt.Sprintf("%x", has)[n : n+size]
	}
	return ""
}
