package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"github.com/chithanh12/kong-plugin/cache"
	"github.com/chithanh12/kong-plugin/utils"
	"regexp"
)

type Config struct {
	WaitTime    int               `json:"waitTime"`
	PublicToken string            `json:"publicToken"`
	Redis       cache.RedisConfig `json:"redis"`
}

var (
	BearerReg = regexp.MustCompile("^Bearer\\s")
)

func (config Config) Access(kong *pdk.PDK) {
	defer func() {
		if r := recover(); r != nil {
			kong.Response.Exit(500, "Gateway Internal Server Error", make(map[string][]string))
		}
	}()

	cookieHeader, _ := kong.Request.GetHeader("cookie")

	cookies := utils.GetCookies(cookieHeader)
	if token, ok := cookies["authorization"]; ok {
		token = BearerReg.ReplaceAllString(token, "")

		authorizationToken := cache.Redis(config.Redis).GetKey(token)
		// write back to response header
		_ = kong.Response.SetHeader("Authorization", authorizationToken)
	}
}

func main() {
	Version := "1.4"
	Priority := 1000
	_ = server.StartServer(func() interface{} { return &Config{} }, Version, Priority)
}
