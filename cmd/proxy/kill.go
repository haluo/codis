package main

import (
	"encoding/json"
	"fmt"

	"github.com/wandoulabs/codis/pkg/proxy"
	"github.com/wandoulabs/codis/pkg/utils/log"
)

type cmdKill struct {
}

func (c *cmdKill) main(d map[string]interface{}) {
	host, auth := d["--admin"].(string), ""
	if s, ok := d["--auth"].(string); ok && s != "" {
		auth = s
	}

	client := proxy.NewApiClient(host)
	info, err := client.GetInfo()
	if err != nil {
		log.PanicErrorf(err, "invalid proxy %s", host)
	}
	client.SetToken(info.Token, auth)

	info.Slots = nil
	info.Stats = nil
	b, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		log.PanicErrorf(err, "json encode failed")
	}
	fmt.Println(string(b))

	if err := client.Shutdown(); err != nil {
		log.PanicErrorf(err, "kill proxy failed")
	}
	fmt.Printf("[KILL] %s\n", host)
}
