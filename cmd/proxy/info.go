package main

import (
	"encoding/json"
	"fmt"

	"github.com/wandoulabs/codis/pkg/proxy"
	"github.com/wandoulabs/codis/pkg/utils/log"
)

type cmdInfo struct {
}

func (c *cmdInfo) main(d map[string]interface{}) {
	host := d["--admin"].(string)

	client := proxy.NewApiClient(host)
	info, err := client.GetInfo()
	if err != nil {
		log.PanicErrorf(err, "invalid proxy %s", host)
	}
	b, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		log.PanicErrorf(err, "json encode error")
	}
	fmt.Println(string(b))
}
