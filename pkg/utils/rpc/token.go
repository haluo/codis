package rpc

import (
	"crypto/md5"
	"fmt"
	"time"
)

func NewToken(initial string) string {
	s := fmt.Sprintf("%s-%d", initial, time.Now().UnixNano())
	b := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", b)
}
