package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123"},
	{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password123"},
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

