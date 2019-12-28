package web

import "github.com/gin-gonic/gin"

// Router returns a http router with endpoints set
func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/products/:id", getProduct)
	router.POST("/products", saveProduct)

	return router
}
