package global

import (
	"github.com/gooderday1805/go-ecommerce-backend-api/pkg/logger"
	"github.com/gooderday1805/go-ecommerce-backend-api/setting"
	"gorm.io/gorm"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	DB     *gorm.DB
	Rdb    *redis.Client
)
