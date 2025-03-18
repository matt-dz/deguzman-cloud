package middleware

import (
	"context"
	"dss/internal/database"
	"dss/internal/logger"
	"net/http"
)

var ctx = context.Background()
var logging = logger.GetLogger()
var db = database.GetDb()

type Middleware func(http.HandlerFunc) http.HandlerFunc

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
