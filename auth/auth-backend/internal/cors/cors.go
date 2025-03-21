package cors

import (
	"net/http"
	"os"
	"regexp"
	"slices"
)

var allowedOrigins = [...]string{"https://auth.deguzman.cloud", "https://deguzman.cloud"}
var localPattern = regexp.MustCompile(`^http:\/\/localhost:\d+$`)

func AddCors(w http.ResponseWriter, r *http.Request) {

	/* If origin is not in allow list, do not add CORS headers */
	origin := r.Header.Get("Origin")
	if os.Getenv("ENV") == "PROD" && !slices.Contains(allowedOrigins[:], origin) {
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", origin)
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Max-Age", "86400")
}

func AddLoginCors(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("ENV") == "PROD" {
		w.Header().Add("Access-Control-Allow-Origin", "auth.deguzman.cloud")
	} else {
		w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	}
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Max-Age", "86400")
}
