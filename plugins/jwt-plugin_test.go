package main

import (
	"testing"

	"github.com/Kong/go-pdk/test"
	"github.com/stretchr/testify/assert"
)

func TestPlugin(t *testing.T) {
	env, err := test.New(t, test.Request{
		Method: "GET",
		Url:    "http://example.com?q=search&x=9",
		Headers: map[string][]string{
			"X-Hi":   {"hello"},
			"cookie": {"Authorization=Bearer xxx;ttt=yyy"},
		},
	})

	assert.NoError(t, err)

	// env.DoHttps(&Config{
	// 	Redis: cache.RedisConfig{
	// 		Host:     "localhost:6379",
	// 		Password: "xbmS3IX9b7",
	// 		DB:       0,
	// 	},
	// })

	assert.Equal(t, "yyyyyyy", env.ClientRes.Headers.Get("authorization"))
}
