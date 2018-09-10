package main

import (
	"flag"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/giuem/ga-proxy/ga"
	"github.com/giuem/ga-proxy/utils"
	"github.com/valyala/fasthttp"
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

func index(ctx *fasthttp.RequestCtx) {
	if len(ctx.Referer()) == 0 || len(ctx.QueryArgs().Peek("ga")) == 0 {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.Response.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Response.Header.Set("Content-Type", "image/gif")

	uid, err := utils.GetOrSetUUID(ctx)
	if err != nil {
		log.Println("[Error] Cannot generate uuid: ", err)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)

	q := ctx.QueryArgs()

	go ga.SendData(&ga.Data{
		UID: uid,
		Tid: q.Peek("ga"),
		Dl:  ctx.Referer(),
		IP:  utils.Realip(ctx),
		Ua:  ctx.UserAgent(),
		Dt:  q.Peek("dt"),
		Dr:  q.Peek("dr"),
		De:  q.Peek("de"),
		Ul:  q.Peek("ul"),
		Sd:  q.Peek("sd"),
		Sr:  q.Peek("sr"),
		Vp:  q.Peek("vp"),
	}, *skipSSLVerify, *debug)
}

func detect(ctx *fasthttp.RequestCtx) {
	ch := make(chan error)
	go func() {
		ch <- ga.Detect(*skipSSLVerify)
	}()
	err := <-ch
	if err != nil {
		log.Println("[Error] detect problem:", err)
		ctx.Response.SetStatusCode(fasthttp.StatusBadGateway)
	} else {
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	}
}

func main() {
	flag.Parse()

	router := fasthttprouter.New()
	router.GET("/", index)
	router.GET("/detect", detect)
	router.HEAD("/detect", detect)

	log.Println("[Info] HTTP server start at: ", *httpAddr)
	log.Fatal(fasthttp.ListenAndServe(*httpAddr, router.Handler))
}
