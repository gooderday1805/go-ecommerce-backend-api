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
	log.Println("🔧 Starting bootstrap process...")

	log.Println("📦 Loading config...")
	LoadConfig()
	log.Println("✅ Config loaded successfully")

	log.Println("📝 Initializing logger...")
	InitLogger()
	log.Println("✅ Logger initialized")

	log.Println("🐬 Connecting to MySQL...")
	InitMysql()
	log.Println("✅ MySQL connected")

	log.Println("🧠 Connecting to Redis...")
	InitRedis()
	log.Println("✅ Redis connected")

	log.Println("🏁 Bootstrap process finished")
}

func startServer() {
	r := InitRouter()

	addr := fmt.Sprintf(":%d", global.Config.Server.Port)
	log.Printf("🚀 Server is running at %s\n", global.Config.Server.Host+addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("❌ Failed to run server: %v", err)
	}
}
