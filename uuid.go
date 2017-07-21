package main

import (
	"net/http"

	"github.com/satori/go.uuid"
)

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
