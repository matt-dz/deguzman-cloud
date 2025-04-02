package middleware

import (
	"crypto/subtle"
	"kube-deploy/internal/cors"
	"kube-deploy/internal/logger"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var log = logger.GetLogger()
var masterKey string

type Middleware func(http.HandlerFunc) http.HandlerFunc

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func init() {
	masterKey = os.Getenv("MASTER_KEY")
	if masterKey == "" {
		panic("MASTER_KEY not set")
	}
}

func (lrw *logResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LogContext() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(logger.AppendCtx(r.Context(), slog.String("method", r.Method)))
			r = r.WithContext(logger.AppendCtx(r.Context(), slog.String("path", r.URL.Path)))
			next(w, r)
		}
	}
}

func Timer() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			lrw := &logResponseWriter{w, http.StatusOK}
			next(lrw, r)
			r = r.WithContext(logger.AppendCtx(r.Context(), slog.Int("status", lrw.statusCode)))
			log.InfoContext(r.Context(), "Request handled", slog.String("duration", time.Since(start).String()))
		}
	}
}

func AuthenticateSecret() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-Api-Key")
			if subtle.ConstantTimeCompare([]byte(apiKey), []byte(masterKey)) != 1 {
				log.ErrorContext(r.Context(), "Unauthorized request", slog.String("error", "invalid API key"))
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next(w, r)
		}
	}
}

func AddCors() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			cors.AddCors(w, r)
			next(w, r)
		}
	}
}

/*
Chain adds middleware in a chained fashion to the HTTP handler.
The middleware is applied in the order in which it is passed.
*/
func Chain(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	// Applied in reverse to preserve the order
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h.ServeHTTP)
	}

	return h
}
