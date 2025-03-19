package middleware

import (
	"context"
	"deguzman-auth/internal/session"
	"errors"
	"log/slog"
	"net/http"
	"os"
)

func ValidateOrigin() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			/* Verify same origin to prevent CSRF attacks */
			if r.Method != "GET" && r.Method != "" {
				origin := r.Header.Get("Origin")
				if origin == "" || origin != os.Getenv("BASE_URL") {
					log.Error("Request not allowed from origin", slog.String("origin", origin))
					http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
					return
				}
			}
		}
	}
}

func HandleRequest() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			userSession, err := session.AuthorizeSession(w, r)
			if err != nil {
				if errors.Is(err, session.ErrUnauthorized) {
					log.Error("Unauthorized request", slog.String("error", err.Error()))
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				} else {
					log.Error("Unable to authorize request", slog.String("error", err.Error()))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}
			r = r.WithContext(context.WithValue(r.Context(), "emailVerified", userSession.EmailVerified))
			next(w, r)
		}
	}
}
