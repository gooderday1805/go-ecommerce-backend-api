package initialize

import (
	"github.com/gin-gonic/gin"
	Controller "github.com/gooderday1805/go-ecommerce-backend-api/internal/controller"
)

func InitRouter() *gin.Engine {
	userController := Controller.NewUserController()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		// User routes
		v1.GET("/users/:id", userController.GetUserById)
		v1.POST("/register", userController.Register)

		// Product routes
		v1.GET("/productsproducts", Controller.GetProductById)
	}

	return r
}