package main

import (
	"net/http"
	"qr-generator/internal/api"
	"qr-generator/internal/logger"
)

var log = logger.GetLogger()

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/qr", api.HandleGenerateCode)
	mux.HandleFunc("GET /api/qr/{id}", api.HandleGetCode)

	log.Info("Starting server on :80")
	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
