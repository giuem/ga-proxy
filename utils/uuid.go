package utils

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/valyala/fasthttp"
)

// GetOrSetUUID get uuid from cookie, or set new
func GetOrSetUUID(ctx *fasthttp.RequestCtx) (string, error) {
	uid := string(ctx.Request.Header.Cookie("uuid"))
	if uid != "" {
		return uid, nil
	}

	ns, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	uid = uuid.NewV5(ns, string(ctx.UserAgent())).String()
	cookie := fasthttp.AcquireCookie()
	cookie.SetKey("uuid")
	cookie.SetValue(uid)
	cookie.SetPath("/")
	cookie.SetExpire(time.Now().AddDate(24, 10, 10))
	cookie.SetHTTPOnly(true)
	ctx.Response.Header.SetCookie(cookie)

	return uid, nil
}
