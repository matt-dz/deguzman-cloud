package middleware

import (
	"net/http"
	"strings"
)

func Authenticate() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logging.Info("Authenticating user")

			/* Retrieving authorization header */
			logging.Info("Retrieving authorization header")
			bearerToken := r.Header.Get("Authorization")
			if bearerToken == "" {
				logging.Error("Authorization header not set")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			/* Extracting bearer token */
			logging.Info("Extracting bearer token")
			token := strings.Split(bearerToken, "Bearer ")
			if len(token) != 2 {
				logging.Error("Bearer token is malformed")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			/* Checking DB for token */
			logging.Info("Checking DB for token")
			var count int
			query := `
			WITH count AS (
				SELECT * FROM tokens WHERE token = $1
			) SELECT COUNT(*) FROM count
			`
			if err := db.QueryRow(ctx, query, token[1]).Scan(&count); err != nil {
				logging.Error("Failed to query db", "error", err.Error())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if count == 0 {
				logging.Error("Token not found")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			logging.Info("User authenticated")
			f(w, r)
		}
	}
}
