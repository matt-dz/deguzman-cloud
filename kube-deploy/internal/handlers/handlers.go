package handlers

import (
	"encoding/json"
	"kube-deploy/internal/cors"
	"kube-deploy/internal/k8s"
	"kube-deploy/internal/logger"
	"log/slog"
	"net/http"
)

var log = logger.GetLogger()

func HandleCors(w http.ResponseWriter, r *http.Request) {
	cors.AddCors(w, r)
}

func HandleGetNamespace(w http.ResponseWriter, r *http.Request) {
	namespace := r.PathValue("namespace")
	exists, err := k8s.NamespaceExists(namespace)
	if err != nil {
		log.ErrorContext(r.Context(), "Unable to retrieve namespace", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

func HandleGetDeploymentLogs(w http.ResponseWriter, r *http.Request) {
	namespace, deploymentName := r.PathValue("namespace"), r.PathValue("name")

	log.DebugContext(r.Context(), "Retrieving Logs")
	logs, err := k8s.GetPodLogs(namespace, deploymentName)
	if err != nil {
		log.ErrorContext(r.Context(), "Failed to retrieve logs", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.DebugContext(r.Context(), "Encoding result")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
