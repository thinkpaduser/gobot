package utils

import (
	"net/http"
	url "net/url"
)

func NewTorClient() *http.Client {
	addr, _ := url.Parse("socks5://127.0.0.1:9050")
	return &http.Client{
		Transport:  &http.Transport{
			Proxy: http.ProxyURL(addr),
		},
	}
}
