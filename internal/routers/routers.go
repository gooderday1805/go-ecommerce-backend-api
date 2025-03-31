package routers

// Trong GO muốn package khác muốn gọi thì function phải viết hoa chữ đầu
import (
	"github.com/gin-gonic/gin"
	"github.com/gooderday1805/go-ecommerce-backend-api/internal/controller/user"
)

func NewRouters() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", user.GetUserById)
		// v1.POST("/users", createUser)
		// v1.GET("/products", getProducts)
		// v1.POST("/products", createProduct)
	}

	return r
}

func getUsers(c *gin.Context) {
	// handler function for GET /api/v1/users
	c.JSON(200, gin.H{"message": "get users"})
}