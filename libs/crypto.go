package libs

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
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

// 唯一哈希(随机)
func (crypto) Unique(size uint8) string {
	return Crypto.MD5(strconv.FormatInt(time.Now().UnixNano()+rand.Int63(), 10), size)
}
