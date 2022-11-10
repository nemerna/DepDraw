package test

import (
	"github.com/RHEcosystemAppEng/DepDraw/src/k8sclient/pkg/config"
	log "github.com/RHEcosystemAppEng/DepDraw/src/k8sclient/pkg/log"
)

func InitTest() *config.Config {
	config := config.NewConfig()
	log.InitLogger(config.LogLevel())
	return config
}
