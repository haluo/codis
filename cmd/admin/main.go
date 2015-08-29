// Copyright 2014 Wandoujia Inc. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package main

import (
	"github.com/docopt/docopt-go"

	"github.com/wandoulabs/codis/pkg/utils/log"
)

func main() {
	const usage = `
Usage:
	codis-admin info --proxy=ADDR
	codis-admin kill --proxy=ADDR [--auth=AUTH]
`

	d, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		log.PanicError(err, "parse arguments failed")
	}

	switch {
	case d["info"]:
		new(cmdInfo).main(d)
	case d["kill"]:
		new(cmdKill).main(d)
	}
}
