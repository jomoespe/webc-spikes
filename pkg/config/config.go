package config

import (
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
