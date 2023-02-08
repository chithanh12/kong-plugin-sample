package utils

import (
	"net/http"
	"strings"
)

func GetCookies(cookieHeaderValue string) map[string]string {
	if cookieHeaderValue == "" {
		return map[string]string{}
	}

	header := http.Header{}
	header.Add("Cookie", cookieHeaderValue)
	request := http.Request{Header: header}
	rs := map[string]string{}
	for _, c := range request.Cookies() {
		rs[strings.ToLower(c.Name)] = c.Value
	}

	return rs
}
