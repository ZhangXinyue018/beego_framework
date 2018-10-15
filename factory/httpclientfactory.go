package factory

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var (
	HttpClient *http.Client
)

func InitClient() {
	HttpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     2 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: 2 * time.Second,
	}
}

func GetHttpClient() (*http.Client) {
	if HttpClient == nil {
		fmt.Println("-----init client------")
		InitClient()
	}
	return HttpClient
}
