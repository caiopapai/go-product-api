package web

import "github.com/gin-gonic/gin"

// Router returns a http router with endpoints set
func Router() *gin.Engine {

	router := gin.Default()

	router.POST("/products", saveProduct)
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.GET("/products/:barcode", getProductByBarcode)
	router.GET("/products/:name", getProductByName)
	router.DELETE("/products/:id", deleteByID)
	router.PUT("/products", updateProduct)

	return router
}
