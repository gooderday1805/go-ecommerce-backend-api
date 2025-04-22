package initialize

import (
	"fmt"
	"log"

	"github.com/gooderday1805/go-ecommerce-backend-api/global"
)

func Run() {
	bootstrap()
	startServer()
}

func bootstrap() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
}

func startServer() {
	r := InitRouter()

	addr := fmt.Sprintf(":%d", global.Config.Server.Port)
	log.Printf("🚀 Server is running at %s\n", global.Config.Server.Host+addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("❌ Failed to run server: %v", err)
	}
}
