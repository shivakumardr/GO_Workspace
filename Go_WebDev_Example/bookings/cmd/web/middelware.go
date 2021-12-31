package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)
/*
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}*/

// NoSurf adds the protection to all POST requests
func NoSurf(next http.Handler)http.Handler{
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad load and save s the session on every requests
func SessionLoad(next http.Handler)http.Handler{
	return session.LoadAndSave(next)
}