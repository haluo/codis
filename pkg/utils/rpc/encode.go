package rpc

import (
	"encoding/base64"

	"github.com/wandoulabs/codis/pkg/utils/errors"
)

func Encode64(addr string) string {
	return base64.StdEncoding.EncodeToString([]byte("#" + addr))
}

func Decode64(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil || len(b) < 1 || b[0] != '#' {
		return "", errors.Errorf("invalid base64 encoding")
	}
	return string(b[1:]), nil
}
