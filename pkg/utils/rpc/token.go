package rpc

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

func NewToken() string {
	hostname, _ := os.Hostname()
	c := make([]byte, 16)
	rand.Read(c)

	s := fmt.Sprintf("%s-%d-%x", hostname, time.Now().UnixNano(), c)
	b := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", b)
}
