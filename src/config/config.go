package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Settings struct {
	resolver string
	start    time.Duration
	refresh  time.Duration
	verbose  bool
	domains  []string
}

const envPrefix = "DCR."

func GetConfig() (*Settings, error) {
	env := GetEnv()
	if len(env) > 0 {
		reso, resoEex := GetValue(env, EnvName("resolver"))
		inte, _ := GetValue(env, EnvName("refresh"))
		dur, durErr := time.ParseDuration(inte)
		insd, _ := GetValue(env, EnvName("start"))
		sd, sdErr := time.ParseDuration(insd)
		verb, verbEex := GetValue(env, EnvName("verbose"))
		doma, domaEex := GetValues(env, EnvName("domains"))
		if resoEex && domaEex && durErr == nil && sdErr == nil {
			verbI := verbEex
			if verbEex {
				bconf, bconfErr := strconv.ParseBool(verb)
				if bconfErr != nil {
					verbI = bconf
				}
			}
			res := &Settings{
				reso,
				sd,
				dur,
				verbI,
				doma,
			}
			if res.verbose {
				fmt.Println("Resolver:", res.resolver)
				fmt.Println("Start delay", res.start)
				fmt.Println("Refresh delay:", res.refresh)
				fmt.Println("Daomains:")
				for _, d := range res.domains {
					fmt.Println("-", d)
				}
			}
			return res, nil
		}
	}
	return nil, errors.New("Environment not set correctly")
}

func (c *Settings) Resolver() string {
	return c.resolver
}

func (c *Settings) Start() time.Duration {
	return c.start
}

func (c *Settings) Refresh() time.Duration {
	return c.refresh
}

func (c *Settings) Verbose() bool {
	return c.verbose
}

func (c *Settings) Domains() []string {
	return c.domains
}

func (c *Settings) VPrintln(text string) {
	if c.Verbose() {
		fmt.Println(text)
	}
}

func (c *Settings) VLine() {
	c.VPrintln("---------------------")
}

func GetEnv() []string {
	completeEnv := os.Environ()
	filteredEnv := []string{}

	for i := range completeEnv {
		if strings.HasPrefix(completeEnv[i], envPrefix) {
			filteredEnv = append(filteredEnv, completeEnv[i])
		}
	}

	return filteredEnv
}

func EnvName(name string) string {
	return envPrefix + name
}

func GetValue(s []string, e string) (string, bool) {
	lowE := strings.ToLower(e)
	for i := range s {
		siSep := strings.Split(s[i], "=")
		if strings.ToLower(siSep[0]) == lowE {
			return strings.TrimSpace(siSep[1]), true
		}
	}
	return "", false
}

func GetValues(s []string, e string) ([]string, bool) {
	lowEPref := strings.ToLower(e) + "["
	res1 := []string{}
	res2 := false
	for i := range s {
		siSep := strings.Split(s[i], "=")
		if strings.HasPrefix(strings.ToLower(siSep[0]), lowEPref) {
			adS := strings.TrimSpace(siSep[1])
			if len(adS) > 0 {
				res1 = append(res1, adS)
				res2 = true
			}
		}
	}
	return res1, res2
}
