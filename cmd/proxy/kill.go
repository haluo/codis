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
	host := d["--admin"].(string)

	force := d["--force"].(bool)

	client := proxy.NewApiClient(host)
	info, err := client.GetInfo()
	if err != nil {
		log.PanicErrorf(err, "invalid proxy %s", host)
	}

	if force {
		err := client.Shutdown(info.Token)
		if err != nil {
			log.PanicErrorf(err, "kill proxy failed")
		}
		fmt.Printf("proxy %s killed\n", host)
	} else {
		info.Stats = nil
		info.Slots = nil
		b, err := json.MarshalIndent(info, "", "    ")
		if err != nil {
			log.PanicErrorf(err, "json encode error")
		}
		fmt.Println(string(b))
	}
}
