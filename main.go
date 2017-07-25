package main

import (
	"log"
	"net/http"

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
}

// ServerHandle handle proxy logic
func ServerHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" ||
		len(req.Referer()) == 0 || len(req.FormValue("ga")) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Content-Type", "image/gif")

	uid := GetOrSetUUID(w, req)

	w.WriteHeader(http.StatusOK)

	go SendData(uid, req)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", ServerHandle)
	log.Println("http server start at:", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
