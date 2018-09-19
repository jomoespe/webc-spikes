// Copyright 2010 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package config

import (
	"regexp"
	"time"
)

type Config struct {
	Name string
	RouteConf
	Routes map[string]Route
}

type RouteConf struct {
	Timeout            duration
	MaxConns           int
	InsecureSkipVerify bool `toml:"insecure-skip-verify"`
	Excepts            []string
	FilterHeaders      []string
	FilterCookies      []string
}
type Route struct {
	To string
	RouteConf
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return
}

var (
	defaultName               = ""
	defaultMaxConns           = 0
	defaultInsecureSkipVerify = false
	defaultTimeout            = duration{time.Duration(30 * time.Second)}
	defaultExcepts            = []string{"/health", "/ready"}
	defaultFilterHeaders      = []string{}
	defaultFilterCookies      = []string{}
)

// NewConfig creates a new configuration with default values.
func NewConfig() (config Config) {
	config.Name = defaultName
	config.Timeout = defaultTimeout
	config.MaxConns = defaultMaxConns
	config.InsecureSkipVerify = defaultInsecureSkipVerify
	config.Excepts = defaultExcepts
	config.FilterHeaders = defaultFilterHeaders
	config.FilterCookies = defaultFilterCookies
	return
}

// ForURI returns the router configuration for an uri
func (c *Config) ForURI(uri string) (route Route, found bool) {
	for k, v := range c.Routes {
		expr, _ := regexp.Compile(k)
		if f := expr.FindStringIndex(uri); f != nil && f[0] == 0 {
			found, route = true, v
			break
		}
	}
	return
}
