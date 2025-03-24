package main

import (
	"deguzman-auth/internal/handlers"
	"deguzman-auth/internal/logger"
	"deguzman-auth/internal/middleware"
	"net/http"
)

var log = logger.GetLogger()

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", middleware.Chain(
		handlers.HandleHeartbeat,
		middleware.LogContext(),
		middleware.Timer(),
	))

	mux.HandleFunc("OPTIONS /", middleware.Chain(
		handlers.HandleCors,
		middleware.LogContext(),
		middleware.Timer(),
	))

	mux.HandleFunc("POST /api/login", middleware.Chain(
		handlers.HandleLogin,
		middleware.LogContext(),
		middleware.ValidateOrigin(),
		middleware.AddLoginCors(),
		middleware.Timer(),
	))

	mux.HandleFunc("POST /api/signup", middleware.Chain(
		handlers.HandleSignup,
		middleware.HandleRequest(),
		middleware.ValidateOrigin(),
		middleware.LogContext(),
		middleware.Timer(),
	))

	mux.HandleFunc("POST /api/logout", middleware.Chain(
		handlers.HandleLogout,
		middleware.HandleRequest(),
		middleware.ValidateOrigin(),
		middleware.LogContext(),
		middleware.AddCors(),
		middleware.Timer(),
	))

	mux.HandleFunc("POST /api/auth", middleware.Chain(
		handlers.HandleAuth,
		middleware.HandleRequest(),
		middleware.ValidateOrigin(),
		middleware.LogContext(),
		middleware.AddCors(),
		middleware.Timer(),
	))

	mux.HandleFunc("POST /api/verify-email", middleware.Chain(
		handlers.HandleEmailVerification,
		middleware.HandleRequest(),
		middleware.ValidateOrigin(),
		middleware.LogContext(),
		middleware.AddCors(),
		middleware.Timer(),
	))

	mux.HandleFunc("POST /api/verify-email/send", middleware.Chain(
		handlers.HandleEmailVerificationSend,
		middleware.HandleRequest(),
		middleware.ValidateOrigin(),
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
