package ga

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/tomasen/realip"
)

const url = "https://www.google-analytics.com/collect"

// SendData send data to google server
func SendData(uid string, req *http.Request, skipSSLVerify, debug bool) {
	newReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("[Error] Cannot make request: ", err)
		return
	}

	newReq.Close = true

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

	tr := &http.Transport{}
	if skipSSLVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(newReq)
	if err != nil {
		log.Println("[Error] Cannot make request: ", err)
		return
	}

	defer resp.Body.Close()
	if debug {
		log.Printf("[Debug] SEND %s\n", newReq.URL.RawQuery)
	}
}

// Detect detect connection with Google
func Detect(skipSSLVerify bool) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Close = true
	tr := &http.Transport{}
	if skipSSLVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
