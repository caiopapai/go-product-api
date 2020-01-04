package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/caiopapai/go-product-api/cmd/domain/entity"

	"github.com/caiopapai/go-product-api/cmd/domain/service"

	"github.com/gin-gonic/gin"
)

//Index redirect to a Index page
func Index(c *gin.Context) {
	service := service.Init()
	products := service.GetAll()
	c.HTML(http.StatusOK, "Index", products)
}

//NewProduct redirects to new product page
func NewProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "NewProduct", nil)
}

//InsertProduct insert new product
func InsertProduct(c *gin.Context) {
	if c.Request.Method == "POST" {

		name := c.PostForm("name")
		barcode := c.PostForm("barcode")
		brandname := c.PostForm("brandname")
		isvegan := c.PostForm("isvegan")
		price := c.PostForm("price")
		duedate := c.PostForm("duedate")
		mescode := c.PostForm("mescode")
		mesvalue := c.PostForm("mesvalue")

		convIsVegan, err := strconv.ParseBool(isvegan)
		if err != nil {
			log.Println("Error to parse Is Vegan:", err)
		}

		p := entity.Product{
			Name:        name,
			Barcode:     barcode,
			BrandName:   brandname,
			DueDate:     duedate,
			IsVegan:     convIsVegan,
			Price:       price,
			Measurement: entity.Measurement{Code: mescode, Value: mesvalue},
		}

		service := service.Init()
		service.Save(&p)
	}

	c.Redirect(http.StatusMovedPermanently, "/")

}

//Remove one product
func Remove(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	service := service.Init()
	service.Remove(productID)
	c.Redirect(http.StatusMovedPermanently, "/")
}
