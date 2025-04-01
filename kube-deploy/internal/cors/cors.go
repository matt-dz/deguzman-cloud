package cors

import "net/http"

func AddCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "X-Api-Key")
	w.Header().Add("Access-Control-Max-Age", "86400")
}
