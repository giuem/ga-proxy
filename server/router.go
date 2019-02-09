package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giuem/ga-proxy/ga"
)

func handlePageView(c *gin.Context) {
	if len(c.Request.Referer()) == 0 || len(c.Query("ga")) == 0 {
		handleRedirect(c)
		return
	}

	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	uid := getUUID(c)

	data := ga.CommonData{
		Version:          1,
		TrackingID:       c.Query("ga"),
		ClientID:         uid,
		UserIP:           c.ClientIP(),
		UserAgent:        c.Request.UserAgent(),
		DocumentReferer:  c.Query("dr"),
		ScreenResolution: c.Query("sr"),
		ViewportSize:     c.Query("vp"),
		DocumentEncoding: c.Query("de"),
		ScreenColors:     c.Query("sd"),
		UserLanguage:     c.Query("ul"),
		DocumentLink:     c.Request.Referer(),
		DocumentTitle:    c.Query("dt"),
	}
	go ga.PageView(data)

	c.Status(http.StatusOK)
}

func handlePing(c *gin.Context) {
	err := ga.Detect()
	if err != nil {
		if c.Request.Method == http.MethodHead {
			c.Status(http.StatusBadGateway)
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"msg": err.Error()})
		}
		return
	}

	if c.Request.Method == http.MethodHead {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	}
}

func handleRedirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://github.com/giuem/ga-proxy")
}
