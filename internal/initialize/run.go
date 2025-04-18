package initialize

import (
	"fmt"
	"github.com/gooderday1805/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	// router
	r := InitRouter()
	r.Run(fmt.Sprintf("Router is running on: %d", global.Config.Server.Port))
}
