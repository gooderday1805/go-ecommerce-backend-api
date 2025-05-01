package initialize

import (
	"log"

	"github.com/gooderday1805/go-ecommerce-backend-api/global"
	"github.com/gooderday1805/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	zapLogger, err := logger.NewLogger(global.Config.Logger.Level, global.Config.Logger.Env == "dev")
	if err != nil {
		log.Fatalf("Initialization log fail: %v", err)
	}

	global.Logger = zapLogger
}
