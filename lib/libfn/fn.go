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

// 这里面的END是为了快速终止运行，避免大量if else或者重复判断
func init() {
	// 自动获取env
	if os.Getenv("http_proxy") != "" {
		// 解析Proxy
		ProxyURL, _ := url.Parse(os.Getenv("http_proxy"))
		// 只解析Proxy的客户端
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(ProxyURL),
			},
		}
		goto END
	}

	// 默认客户端
	httpClient = http.DefaultClient
END:
}

func Client() *http.Client {
	return httpClient
}
