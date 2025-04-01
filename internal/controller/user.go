package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gooderday1805/go-ecommerce-backend-api/internal/service"
	"net/http"
)
type UserController struct {
	UserService	*service.UserService 
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, uc.UserService.GetInforUsers())
}
