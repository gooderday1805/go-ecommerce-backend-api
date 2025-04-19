	package initialize

	import (
		"fmt"
		"github.com/spf13/viper"

		"github.com/gooderday1805/go-ecommerce-backend-api/global"
	)

	func LoadConfig() {
		vipper := viper.New()
		vipper.SetConfigName("local")
		vipper.SetConfigType("yaml")
		vipper.AddConfigPath("./config/")

		// read config
		err := vipper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Failed to read config %w", err))
		}

		fmt.Println("Server Port::: ", vipper.GetInt("server.port"))

		// Unmarshal -> data
		err = vipper.Unmarshal(&global.Config)
		if err != nil {
			panic(fmt.Errorf("Failed to unmarshal config %w", err))
		}
	}
