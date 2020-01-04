package web

import (
	"encoding/json"
	"net/http"

	"github.com/caiopapai/go-product-api/cmd/api"
	"github.com/caiopapai/go-product-api/cmd/api/response"
	"github.com/caiopapai/go-product-api/cmd/domain/service"
	"github.com/gin-gonic/gin"
)

func saveProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	service := service.Init()

	var metadata api.Metadata

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := []byte(buf[0:num])

	json.Unmarshal(reqBody, &metadata)

	product := metadata.Data.Product
	productID := service.Save(&product)
	response := response.Metadata{}
	response.Data.ProductID = productID
	response.Data.Product = product

	c.JSON(http.StatusOK, response)
}

//get all products
func getProducts(c *gin.Context) {
	service := service.Init()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, service.GetAll())
}

func getProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(http.StatusNotImplemented, "TODO")
}

func getProductByBarcode(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(http.StatusNotImplemented, "TODO")
}

func getProductByName(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(http.StatusNotImplemented, "TODO")
}

func remove(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, "TODO")
}

func updateProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(http.StatusNotImplemented, "TODO")
}
