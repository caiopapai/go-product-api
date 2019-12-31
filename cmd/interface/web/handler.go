package web

import (
	controller "github.com/caiopapai/go-product-api/cmd/controller"
	"github.com/gin-gonic/gin"
)

// Router returns a http router with endpoints set
func Router() *gin.Engine {

	router := gin.Default()
	//To serve MVC
	router.LoadHTMLGlob("interface/template/*")
	router.GET("/", controller.Index)
	router.GET("/new", controller.NewProduct)
	router.POST("/insert", controller.InsertProduct)

	//To serve API
	router.POST("/products", saveProduct)
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.GET("/products/:barcode", getProductByBarcode)
	router.GET("/products/:name", getProductByName)
	router.DELETE("/products/:id", deleteByID)
	router.PUT("/products", updateProduct)
	return router
}
