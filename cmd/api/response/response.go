package response

import "github.com/caiopapai/go-product-api/cmd/domain/entity"

//Metadata is the struct that holds the Object
type Metadata struct {
	Data Data `json:"data"`
}

//Data is the struct that holds the Object
type Data struct {
	ProductID int            `json:"id"`
	Product   entity.Product `json:"product"`
}
