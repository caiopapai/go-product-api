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
	router.GET("/delete", controller.Remove)
	router.POST("/insert", controller.InsertProduct)
	router.GET("/edit", controller.Edit)
	router.POST("/update", controller.Update)

	//To serve API
	router.POST("/products", saveProduct)
	router.GET("/products", getProducts)
	router.GET("/products/:name", getProductByName)
	router.DELETE("/products/:id", remove)
	router.PUT("/products", updateProduct)
	return router
}
