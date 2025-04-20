package global

import (
	"github.com/gooderday1805/go-ecommerce-backend-api/setting"
	"github.com/gooderday1805/go-ecommerce-backend-api/pkg/logger"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)