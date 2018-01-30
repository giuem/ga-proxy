package utils

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/valyala/fasthttp"
)

// GetOrSetUUID get uuid from cookie, or set new
func GetOrSetUUID(ctx *fasthttp.RequestCtx) ([]byte, error) {
	uid := ctx.Request.Header.Cookie("uuid")
	if len(uid) > 0 {
		return uid, nil
	}

	ns, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	uid = uuid.NewV5(ns, string(ctx.UserAgent())).Bytes()
	cookie := fasthttp.AcquireCookie()
	cookie.SetKey("uuid")
	cookie.SetValueBytes(uid)
	cookie.SetPath("/")
	cookie.SetExpire(time.Now().AddDate(24, 10, 10))
	cookie.SetHTTPOnly(true)
	ctx.Response.Header.SetCookie(cookie)

	return uid, nil
}
