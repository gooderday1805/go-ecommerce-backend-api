package global

import (
	"github.com/gooderday1805/go-ecommerce-backend-api/pkg/logger"
	"github.com/gooderday1805/go-ecommerce-backend-api/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	DB     *gorm.DB
)
