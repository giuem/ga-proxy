package main

import (
	"log"
	"net/http"

	"flag"

	"github.com/satori/go.uuid"
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
	client := &http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	log.Printf("SEND %s\n", newReq.URL.RawQuery)
}

// GetOrSetUUID get uuid of current user or set new
func GetOrSetUUID(w http.ResponseWriter, req *http.Request) string {
	cookie, err := req.Cookie("uuid")
	var uid string
	if err == http.ErrNoCookie {
		ns := uuid.NewV4()
		uid = uuid.NewV5(ns, req.Form.Encode()+req.UserAgent()+req.RemoteAddr).String()
		http.SetCookie(w, &http.Cookie{
			Name:     "uuid",
			Value:    uid,
			Path:     "/",
			MaxAge:   315360000,
			HttpOnly: true,
		})
	} else {
		uid = cookie.Value
	}
	return uid
}

// ServerHandle handle proxy logic
func ServerHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}

	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Content-Type", "image/gif")

	uid := GetOrSetUUID(w, req)

	w.WriteHeader(http.StatusNoContent)

	go SendData(uid, req)
}

func init() {
	http.HandleFunc("/", ServerHandle)
}

func main() {
	httpAddr := flag.String("http_addr", ":80", "http listen addr")
	flag.Parse()

	log.Println("http server start at:", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
