package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HealthCheck godoc
// @Summary Health check
// @Description Health check if server is ready
// @Tags health
// @Success 204 "Server is ready"
// @Router /healthz [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Info("HealthCheckHandler is called")
	w.WriteHeader(http.StatusNoContent)
}
