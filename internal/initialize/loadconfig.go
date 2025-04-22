package initialize

import (
	"fmt"

	"github.com/gooderday1805/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	vipper := viper.New()
	vipper.SetConfigName("local")
	vipper.SetConfigType("yaml")
	vipper.AddConfigPath("./config/")

	// Read
	err := vipper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("❌ Failed to read config %w", err))
	}

	// Unmarshal -> data
	err = vipper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("❌ Failed to unmarshal config %w", err))
	}
}
