package api

import (
	"github.com/caiopapai/go-product-api/cmd/product"
)

//Data is the struct that holds the Object
type Data struct {
	Product product.Product `json:"product"`
}
