package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Run starts a HTTP server
func Run(ip, port string) {
	addr := fmt.Sprintf("%v:%v", ip, port)

	r := gin.New()
	r.Use(gin.Logger())

	r.NoRoute(handleRedirect)

	r.GET("/", handlePageView)
	r.GET("/ping", handlePing)
	r.HEAD("/ping", handlePing)

	r.Run(addr)
}
