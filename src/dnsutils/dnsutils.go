package dnsutils

import (
	"context"
	"dnscacherefresh/config"
	"errors"
	"fmt"
	"net"
	"time"
)

func GetResolverEx(c *config.Settings) (*net.Resolver, error) {
	addr := net.ParseIP(c.Resolver())
	ip := c.Resolver()
	if addr != nil {
		if c.Verbose() {
			fmt.Println("GetResolverEx:", ip)
		}
		return GetResolver(addr.String()), nil
	} else {
		tReso := GetResolver("127.0.0.11")
		dRes, dResErr := LookUp(c, tReso, c.Resolver())
		if dResErr == nil {
			ip = dRes[0]
			if c.Verbose() {
				fmt.Println("GetResolverEx:", c.Resolver(), "(", ip, ")")
			}
			return GetResolver(ip), nil
		}
	}
	return nil, errors.New("Can't get resolver for " + c.Resolver())
}

func GetResolver(resolver string) *net.Resolver {
	res := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, resolver+":53")
		},
	}
	return res
}

func LookUp(conf *config.Settings, resolver *net.Resolver, domain string) ([]string, error) {
	res, resErr := resolver.LookupHost(context.Background(), domain)
	if conf.Verbose() {
		fmt.Println("LookUp:", domain, "Result:", res)
	}
	return res, resErr
}
