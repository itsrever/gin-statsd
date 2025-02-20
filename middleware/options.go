package middleware

import (
	"fmt"
)

// Options holds configurable user settings
type Options struct {
	Host       string
	Port       int
	RequestKey string
	Prefix     string
}

const defaultHost = "127.0.0.1"
const defaultPort = 8125
const defaultRequestKey = "statsdKey"

func (o *Options) getAddress() string {
	host := defaultHost
	port := defaultPort
	if o.Host != "" {
		host = o.Host
	}
	if o.Port > 0 {
		port = o.Port
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func (o *Options) getPrefix() string {
	if o.Prefix != "" {
		return o.Prefix
	}
	return ""
}

func (o *Options) getRequestKey() string {
	if o.RequestKey != "" {
		return o.RequestKey
	}
	return defaultRequestKey
}
