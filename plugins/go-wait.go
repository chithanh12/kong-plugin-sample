package main

import (
	"fmt"
	"time"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

type Config struct {
	WaitTime int
	Secret   string
}

func New() interface{} {
	return &Config{}
}

var requests = make(map[string]time.Time)

func (config Config) Access(kong *pdk.PDK) {
	_ = kong.Response.SetHeader("x-wait-time", fmt.Sprintf("%d seconds", config.WaitTime))

	host, _ := kong.Request.GetHost()
	p, _ := kong.Request.GetPath()
	kong.Log.Debug("Path Request" + p)

	lastRequest, exists := requests[host]
	_ = kong.Response.SetHeader("x-plugin-path", p)

	kong.Log.Debug("Secret Key: " + config.Secret)

	token, _ := kong.Request.GetHeader("Token")
	if token != "" {
		kong.Log.Debug("hello world debug")
	}

	if exists && time.Now().Sub(lastRequest) < time.Duration(config.WaitTime)*time.Second {
		kong.Response.Exit(400, "Maximum Requests Reached", make(map[string][]string))
	} else {
		requests[host] = time.Now()
	}
}

func main() {
	Version := "1.2"
	Priority := 1000
	_ = server.StartServer(New, Version, Priority)
}
