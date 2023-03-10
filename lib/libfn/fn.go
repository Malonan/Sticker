package fn

import (
	"net/http"
	"net/url"
	"os"
)

var (
	httpClient *http.Client
)

type Do struct{}

func init() {
	if os.Getenv("http_proxy") != "" {
		ProxyURL, _ := url.Parse(os.Getenv("http_proxy"))
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(ProxyURL),
			},
		}
	} else {
		httpClient = http.DefaultClient
	}
}

func Client() *http.Client {
	return httpClient
}
