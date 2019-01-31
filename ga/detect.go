package ga

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Detect detects the connectivity with Google Analytics server
func Detect(skipSSLVerify bool) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	// httpClient.TLSConfig.InsecureSkipVerify = skipSSLVerify

	resp := fasthttp.AcquireResponse()

	err := httpClient.Do(req, resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("Return not ok")
	}
	return nil
}
