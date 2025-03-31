package main

import (
	"log"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Tạo router Gin
	r := gin.Default()

	// Thêm CORS middleware để tránh lỗi khi gọi API từ trình duyệt/Postman
	r.Use(cors.Default())

	// Định nghĩa route GET /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Chạy server trên cổng 8080, lắng nghe trên tất cả địa chỉ IP
	port := ":8080"
	log.Printf("Server is running on http://localhost%s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
