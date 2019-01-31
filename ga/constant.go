package ga

import (
	"crypto/tls"

	"github.com/valyala/fasthttp"
)

const url = "https://www.google-analytics.com/collect"

var httpClient = &fasthttp.Client{
	TLSConfig: &tls.Config{
		InsecureSkipVerify: false,
	},
}
