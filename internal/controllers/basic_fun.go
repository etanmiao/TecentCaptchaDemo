package controllers

import (
	"net/http"
	"net/url"
	"os"
)

//SetProxy 设置内网代理，开发网访问
func SetProxy(_ *http.Request) (*url.URL, error) {
	var proxyURL *url.URL
	env := os.Getenv("GO_ENV")
	if env == "dev" {
		return url.Parse("http://127.0.0.1:12639")
	}
	return proxyURL, nil
}
