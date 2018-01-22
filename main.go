package main

import (
	"log"
	"net/http"

	"github.com/giuem/ga-proxy/ga"
	uuid "github.com/satori/go.uuid"

	"flag"
)

var (
	debug         = flag.Bool("debug", false, "output debug info")
	skipSSLVerify = flag.Bool("skip_ssl", false, "skip SSL verify")
	httpAddr      = flag.String("listen", ":80", "listen address")
)

func init() {
	flag.BoolVar(debug, "d", false, "output debug info")
	flag.BoolVar(skipSSLVerify, "s", false, "skip SSL verify")
	flag.StringVar(httpAddr, "l", ":80", "listen address")

	http.HandleFunc("/", serverHandle)
	http.HandleFunc("/detect", serverDetect)
}

func serverHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" ||
		len(req.Referer()) == 0 || len(req.FormValue("ga")) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Content-Type", "image/gif")

	uid, err := getOrSetUUID(w, req)
	if err != nil {
		log.Println("[Error] Cannot generate uuid: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)

	go ga.SendData(uid, req, *skipSSLVerify, *debug)
}

func serverDetect(w http.ResponseWriter, req *http.Request) {
	err := ga.Detect(*skipSSLVerify)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	flag.Parse()
	log.Println("[Info] HTTP server start at: ", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

func getOrSetUUID(w http.ResponseWriter, req *http.Request) (string, error) {
	cookie, err := req.Cookie("uuid")
	var uid string
	if err == http.ErrNoCookie {
		ns, err := uuid.NewV4()
		if err != nil {
			return "", nil
		}

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
	return uid, nil
}
