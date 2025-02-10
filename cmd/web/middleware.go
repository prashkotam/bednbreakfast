package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// prints to console on every request
func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)

	})
}

// It adds/updates CSRF token to a request
func NoSurf(next http.Handler) http.Handler {

	csrfhandler := nosurf.New(next)

	csrfhandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   App.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfhandler
}

// Loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
