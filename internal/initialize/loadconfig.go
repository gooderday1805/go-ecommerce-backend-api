package initialize

import (
	"fmt"
	"log"
	"strings"

	"github.com/gooderday1805/go-ecommerce-backend-api/global"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("⚠️ Warning: Could not load .env file: %v", err)
	}

	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Định dạng cấu hình
	v.SetConfigName("local")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config/")

	// Đọc file cấu hình
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("❌ Failed to read config: %w", err))
	}

	// Unmarshal vào global config
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("❌ Failed to unmarshal config: %w", err))
	}
}
