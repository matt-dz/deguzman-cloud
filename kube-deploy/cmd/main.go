package main

import (
	"kube-deploy/internal/handlers"
	"kube-deploy/internal/logger"
	"kube-deploy/internal/middleware"
	"net/http"
)

var log = logger.GetLogger()

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", middleware.Chain(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.LogContext(),
		middleware.Timer(),
	))

	mux.HandleFunc("OPTIONS /", middleware.Chain(
		handlers.HandleCors,
		middleware.LogContext(),
		middleware.Timer(),
	))

	mux.HandleFunc("GET /namespace/{namespace}", middleware.Chain(
		handlers.HandleGetNamespace,
		middleware.AuthenticateSecret(),
		middleware.LogContext(),
		middleware.AddCors(),
		middleware.Timer(),
	))

	mux.HandleFunc("GET /deployment/{namespace}/{name}/logs", middleware.Chain(
		handlers.HandleGetDeploymentLogs,
		middleware.AuthenticateSecret(),
		middleware.LogContext(),
		middleware.AddCors(),
		middleware.Timer(),
	))

	log.Info("Starting server on :80")
	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
