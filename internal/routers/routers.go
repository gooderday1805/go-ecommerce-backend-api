package routers

// Trong GO muốn package khác muốn gọi thì function phải viết hoa chữ đầu
import (
	"github.com/gin-gonic/gin"
	Controller "github.com/gooderday1805/go-ecommerce-backend-api/internal/controller"
)

func NewRouters() *gin.Engine {
	r := gin.Default()

	// Create a new user controller
	userController := Controller.NewUserController()

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

func getUsers(c *gin.Context) {
	// handler function for GET /api/v1/users
	c.JSON(200, gin.H{"message": "get users"})
}
