package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.New()
	r.Use(gin.Logger())

	r.NoRoute(handleRedirect)

	r.GET("/", handlePageView)
	r.GET("/ping", handlePing)
	r.HEAD("/ping", handlePing)
}

// Run starts a HTTP server
func Run(ip, port string) {
	addr := fmt.Sprintf("%v:%v", ip, port)

	r.Run(addr)
}
