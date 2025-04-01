package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity"`
}

var products = []Product{
	{ID: "1", Name: "Sản phẩm 1", Price: 10000, Quantity: 10},
	{ID: "2", Name: "Sản phẩm 2", Price: 20000, Quantity: 20},
	{ID: "3", Name: "Sản phẩm 3", Price: 30000, Quantity: 30},
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Sản phẩm không tìm thấy"})
}