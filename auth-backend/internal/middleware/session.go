package middleware

import (
	"context"
	"deguzman-auth/internal/session"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"regexp"
)

var (
	originPattern = regexp.MustCompile(`^(?:https:\/\/)?[\w-]*\.?deguzman\.cloud(?::\d{1,5})?$`) // Any origin must be a subdomain of deguzman.cloud
	localPattern  = regexp.MustCompile(`^(?:https?:\/\/)?[\w-]*\.?localhost(?::\d{1,5})?$`)      // Any origin must be a subdomain of localhost
)

func ValidateOrigin() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			/* Verify same origin to prevent CSRF attacks */
			if r.Method != "GET" && r.Method != "" {
				origin := r.Header.Get("Origin")
				if os.Getenv("ENV") == "PROD" && !originPattern.MatchString(origin) {
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
