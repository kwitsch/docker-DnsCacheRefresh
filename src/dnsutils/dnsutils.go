package dnsutils

import (
	"context"
	"dnscacherefresh/config"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

func GetResolverLoop(c *config.Settings) (*net.Resolver, error) {
	for i := 0; i < int(c.Start().Seconds()); i++ {
		r, rErr := GetResolverEx(c)
		if rErr == nil {
			return r, rErr
		} else {
			time.Sleep(time.Second)
		}
	}
	return nil, errors.New("Can't get resolver for " + c.Resolver())
}

func GetResolverEx(c *config.Settings) (*net.Resolver, error) {
	parts := strings.Split(c.Resolver(), ":")
	name := parts[0]
	port := "53"
	if len(parts) > 1 {
		port = parts[1]
	}
	addr := net.ParseIP(name)
	if addr != nil {
		if c.Verbose() {
			fmt.Println("GetResolverEx: ", name+":"+port)
		}
		return GetResolver(addr.String(), port), nil
	} else {
		tReso := GetResolver("127.0.0.11", "53")
		dRes, dResErr := LookUp(c, tReso, c.Resolver())
		if dResErr == nil {
			ip := dRes[0]
			if c.Verbose() {
				fmt.Println("GetResolverEx:", c.Resolver(), "(", ip, ")")
			}
			return GetResolver(ip, port), nil
		}
	}
	return nil, errors.New("Can't get resolver for " + c.Resolver())
}

func GetResolver(name, port string) *net.Resolver {
	res := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, name+":"+port)
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
