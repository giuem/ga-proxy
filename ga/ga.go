package ga

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

const url = "https://www.google-analytics.com/collect"

var httpClient = &fasthttp.Client{
	TLSConfig: &tls.Config{
		InsecureSkipVerify: false,
	},
}

// Data analytics data
type Data struct {
	UID string
	Tid []byte
	Dl  []byte
	IP  string
	Ua  []byte
	Dt  []byte
	Dr  []byte
	Ul  []byte
	Sd  []byte
	Sr  []byte
	Vp  []byte
}

// SendData send data to google server
func SendData(data *Data, skipSSLVerify, debug bool) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	q := req.URI().QueryArgs()
	q.Add("v", "1")
	q.Add("t", "pageview")
	q.AddBytesV("tid", data.Tid)
	q.Add("cid", data.UID)
	q.AddBytesV("dl", data.Dl)
	q.Add("uip", data.IP)
	q.AddBytesV("ua", data.Ua)
	q.AddBytesV("dt", data.Dt)
	q.AddBytesV("dr", data.Dr)
	q.AddBytesV("ul", data.Ul)
	q.AddBytesV("sd", data.Sd)
	q.AddBytesV("sr", data.Sr)
	q.AddBytesV("vp", data.Vp)
	// q.AddBytesV("z", ctx.QueryArgs().Peek("z"))

	httpClient.TLSConfig.InsecureSkipVerify = skipSSLVerify
	// resp := fasthttp.AcquireResponse()
	err := httpClient.Do(req, nil)

	if debug {
		log.Printf("[Debug] SEND %s\n", q.String())
	}

	if err != nil {
		log.Println("[Error] Cannot make request: ", err)
	}
}

// Detect detect connection with Google
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

// // Detect detect connection with Google
// func Detect(skipSSLVerify bool) error {
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return err
// 	}

// 	req.Close = true
// 	tr := &http.Transport{}
// 	if skipSSLVerify {
// 		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
// 	}
// 	client := &http.Client{Transport: tr}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	return nil
// }
