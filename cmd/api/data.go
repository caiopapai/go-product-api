package api

import "github.com/caiopapai/go-product-api/cmd/product"

//Metadata is the struct that holds the Object
type Metadata struct {
	Data Data `json:"data"`
}

//Data is the struct that holds the Object
type Data struct {
	Product product.Product `json:"product"`
}
