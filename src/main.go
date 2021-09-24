package main

import (
	"dnscacherefresh/config"
	"dnscacherefresh/dnsutils"
	"fmt"
	"os"
	"time"
)

func main() {
	c, cErr := config.GetConfig()
	if cErr == nil {
		c.VLine()
		c.VPrintln("Get DNS resolver")
		r, rErr := dnsutils.GetResolverLoop(c)
		c.VLine()
		if rErr == nil {
			for {
				for _, d := range c.Domains() {
					dnsutils.LookUp(c, r, d)
				}
				c.VLine()
				time.Sleep(c.Refresh())
			}
		} else {
			fmt.Println(rErr.Error())
			os.Exit(3)
		}
	} else {
		fmt.Println(cErr.Error())
		os.Exit(2)
	}
	os.Exit(1)
}
