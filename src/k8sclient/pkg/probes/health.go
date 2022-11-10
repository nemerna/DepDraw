package probes

import (
	"encoding/json"
	"net/http"

	"github.com/RHEcosystemAppEng/DepDraw/src/k8sclient/pkg/config"
	log "github.com/RHEcosystemAppEng/DepDraw/src/k8sclient/pkg/log"
)

type HealthCheck struct {
	config *config.Config
}

type HealthStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewHealthCheckForConfig(config *config.Config) *HealthCheck {
	return &HealthCheck{config: config}
}

func (h HealthCheck) Check(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	status := HealthStatus{Status: http.StatusOK, Message: "Hello from me!"}
	jsonResp, err := json.Marshal(status)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	log.Debugf("Health check returns %v", status)
	rw.Write(jsonResp)
}
