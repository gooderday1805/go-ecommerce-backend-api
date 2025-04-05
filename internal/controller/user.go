package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gooderday1805/go-ecommerce-backend-api/internal/models"
	"github.com/gooderday1805/go-ecommerce-backend-api/internal/service"
	"net/http"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, uc.UserService.GetInforUsers())
}

// Register handles the HTTP request for user registration
func (uc *UserController) Register(c *gin.Context) {
	var req models.RegisterRequest

	// Bind the request body to the RegisterRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to register the user
	response, err := uc.UserService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	c.JSON(http.StatusCreated, response)
}
