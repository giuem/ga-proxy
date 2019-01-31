package ga

import (
	"log"

	"github.com/valyala/fasthttp"
)

// PageViewData is analytics data for `t=pageview`
type PageViewData struct {
	UID string
	Tid []byte
	Dl  []byte
	IP  string
	Ua  []byte
	Dt  []byte
	Dr  []byte
	De  []byte
	Ul  []byte
	Sd  []byte
	Sr  []byte
	Vp  []byte
}

// SendPageViewData send data to google server
func SendPageViewData(data *PageViewData, skipSSLVerify, debug bool) {
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
	q.AddBytesV("de", data.De)
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
