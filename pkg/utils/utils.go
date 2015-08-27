// Copyright 2014 Wandoujia Inc. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package utils

import (
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/wandoulabs/codis/pkg/utils/errors"
)

var BootTime = time.Now()

func ResolveGlobalAddr(network string, address string) (string, error) {
	switch network {
	default:
		return "", errors.Errorf("invalid network '%s'", network)
	case "unix", "unixpacket":
		return address, nil
	case "tcp", "tcp4", "tcp6":
		addr, err := net.ResolveTCPAddr(network, address)
		if err != nil {
			return "", errors.Trace(err)
		}
		if ipv4 := addr.IP.To4(); ipv4 != nil {
			if !net.IPv4zero.Equal(ipv4) {
				return addr.String(), nil
			}
		} else if ipv6 := addr.IP.To16(); ipv6 != nil {
			if !net.IPv6zero.Equal(ipv6) {
				return addr.String(), nil
			}
		}
		ifaddrs, err := net.InterfaceAddrs()
		if err != nil {
			return "", errors.Trace(err)
		}
		for _, ifaddr := range ifaddrs {
			switch in := ifaddr.(type) {
			case *net.IPNet:
				if in.IP.IsGlobalUnicast() {
					addr.IP = in.IP
					return addr.String(), nil
				}
			}
		}
		return "", errors.Errorf("no global unicast address is configured")
	}
}

func ValidateProductName(product string) (string, error) {
	name := strings.TrimSpace(product)
	if regexp.MustCompile(`^\w[\w\.\-]*$`).MatchString(name) {
		return name, nil
	} else {
		return "", errors.Errorf("invalid product name, invalid character or bad leading character")
	}
}
