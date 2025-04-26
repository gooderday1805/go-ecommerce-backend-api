// package initialize

// import (
//     "fmt"
//     "log"
// 	"os"
//     "github.com/gooderday1805/go-ecommerce-backend-api/global"
//     "github.com/joho/godotenv"
//     "github.com/spf13/viper"
// )

// func LoadConfig() {
//     // Load .env file first
//     err := godotenv.Load()
//     if err != nil {
//         log.Printf("⚠️ Warning: Could not load .env file: %v", err)
//     }

//     // Create a new Viper instance
//     vipper := viper.New()
//     vipper.SetConfigName("local")      // Đặt tên file cấu hình là "local"
//     vipper.SetConfigType("yaml")      // Định dạng cấu hình là YAML
//     vipper.AddConfigPath("./config/") // Đường dẫn đến thư mục chứa file cấu hình

//     // Tự động ánh xạ các biến môi trường từ hệ thống hoặc từ .env
//     vipper.AutomaticEnv()

//     // Đọc file cấu hình
//     err = vipper.ReadInConfig()
//     if err != nil {
//         panic(fmt.Errorf("❌ Failed to read config: %w", err))
//     }

//     // Giải nén dữ liệu từ cấu hình vào đối tượng global.Config
//     err = vipper.Unmarshal(&global.Config)
//     if err != nil {
//         panic(fmt.Errorf("❌ Failed to unmarshal config: %w", err))
//     }

//     // In mật khẩu để kiểm tra
//     fmt.Println("Password là :", global.Config.Mysql.Password)
// 	fmt.Println("MYSQL_USER:", os.Getenv("MYSQL_USER"))
// fmt.Println("MYSQL_PASSWORD:", os.Getenv("MYSQL_PASSWORD"))
// }
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
	// Đọc file .env
    err := godotenv.Load()
    if err != nil {
        log.Printf("⚠️ Warning: Could not load .env file: %v", err)
    }

    v := viper.New()
	// Tự động ánh xạ các biến môi trường từ hệ thống hoặc từ .env
    v.AutomaticEnv()
    // Chuyển đổi cấu trúc như mysql.password thành MYSQL_PASSWORD
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
    
    fmt.Println("Password là:", global.Config.Mysql.Password)
}