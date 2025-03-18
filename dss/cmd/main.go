package main

import (
	"dss/internal/api"
	"dss/internal/logger"
	"dss/internal/middleware"
	"net/http"
)

var logging = logger.GetLogger()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{key}", middleware.Chain(
		api.HandleGetFile,
		middleware.Authenticate(),
	))

	mux.HandleFunc("GET /list", middleware.Chain(
		api.HandleListFiles,
		middleware.Authenticate(),
	))

	mux.HandleFunc("POST /upload", middleware.Chain(
		api.HandleUploadFile,
		middleware.Authenticate(),
	))

	mux.HandleFunc("GET /heartbeat", middleware.Chain(
		api.HandleHeartbeat,
		middleware.Authenticate(),
	))

	logging.Info("Starting server on :80")
	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
