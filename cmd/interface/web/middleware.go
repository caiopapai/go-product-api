package web

import (
	"encoding/json"
	"net/http"

	"github.com/caiopapai/go-product-api/cmd/api"
	"github.com/caiopapai/go-product-api/cmd/api/response"
	"github.com/caiopapai/go-product-api/cmd/domain/repository"
	"github.com/caiopapai/go-product-api/cmd/domain/service"
	"github.com/caiopapai/go-product-api/cmd/infra/connection"
	"github.com/gin-gonic/gin"
)

func saveProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	repo := repository.ProductRepository{DB: connection.GetConnection()}
	service := service.ProductService{Repository: &repo}

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

func getProducts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(http.StatusNotImplemented, "TODO")
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

func deleteByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(http.StatusNotImplemented, "TODO")
}

func updateProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.String(http.StatusNotImplemented, "TODO")
}
