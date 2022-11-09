package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RHEcosystemAppEng/DepDraw/src/k8sclient/pkg/config"
	log "github.com/RHEcosystemAppEng/DepDraw/src/k8sclient/pkg/log"
	"github.com/RHEcosystemAppEng/DepDraw/src/k8sclient/pkg/probes"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func startApi(config *config.Config) {
	health := probes.NewHealthCheckForConfig(config)
	router.Path("/health").HandlerFunc(health.Check).Name("healthHandler")

	url := fmt.Sprintf(":%d", config.ServerPort())
	log.Infof("Starting listener as %s", url)
	if err := http.ListenAndServe(url, router); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := config.NewConfig()
	log.InitLogger(config.LogLevel())
	log.Infof("Launching %s", os.Args[0])
	log.Infof("Config is %v+", config)
	startApi(config)
}
