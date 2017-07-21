package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/tomasen/realip"
)

// SendData send data to google server
func SendData(uid string, req *http.Request) {
	newReq, err := http.NewRequest("GET", "https://www.google-analytics.com/collect", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := newReq.URL.Query()
	q.Add("v", "1")
	q.Add("t", "pageview")
	q.Add("tid", req.FormValue("ga"))
	q.Add("cid", uid)
	q.Add("dl", req.Referer())
	q.Add("uip", realip.RealIP(req))
	q.Add("ua", req.UserAgent())
	q.Add("dt", req.FormValue("dt"))
	q.Add("dr", req.FormValue("dr"))
	q.Add("ul", req.FormValue("ul"))
	q.Add("sd", req.FormValue("sd"))
	q.Add("sr", req.FormValue("sr"))
	q.Add("vp", req.FormValue("vp"))
	q.Add("z", req.FormValue("z"))
	newReq.URL.RawQuery = q.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(newReq)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if *debug {
		log.Printf("SEND %s\n", newReq.URL.RawQuery)
	}
}
