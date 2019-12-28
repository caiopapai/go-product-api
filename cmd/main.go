package main

import (
	"net/http"

	"github.com/caiopapai/go-product-api/cmd/domain/entity"
	"github.com/caiopapai/go-product-api/cmd/infra/connection"

	"github.com/caiopapai/go-product-api/cmd/domain/repository"
	"github.com/caiopapai/go-product-api/cmd/domain/service"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	repo := repository.ProductRepository{DB: connection.GetConnection()}
	service := service.ProductService{Repository: &repo}

	router.GET("/product", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, productFactory())
	})

	router.POST("/product", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var product entity.Product
		c.BindJSON(&product)
		//json.Unmarshal(posted, &product)
		c.JSON(http.StatusOK, service.Save(&product))
	})

	router.Run(":8080")
}

func productFactory() *entity.Product {

	p := entity.Product{}
	p.Name = "Lisoform"
	p.DueDate = "2019-07-01"
	p.IsVegan = false
	p.Barcode = "784w8947847927482374t"
	p.BrandName = "Xé"
	p.Price = "8.90"
	p.Measurement.Code = "400"
	p.Measurement.Value = "333"

	return &p
}
