package initialize

import (
	// "fmt"

	// "github.com/gooderday1805/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	InitRouter()
}