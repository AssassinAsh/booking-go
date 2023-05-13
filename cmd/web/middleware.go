package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	noSurf := nosurf.New(next)
	noSurf.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})
	return noSurf
}

func SetSession(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
