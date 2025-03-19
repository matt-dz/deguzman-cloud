package middleware

import (
	"deguzman-auth/internal/logger"
	"log/slog"
	"net/http"
	"time"
)

var log = logger.GetLogger()

type Middleware func(http.HandlerFunc) http.HandlerFunc

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
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
