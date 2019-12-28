package service

import (
	"github.com/caiopapai/go-product-api/cmd/domain/repository"

	"github.com/caiopapai/go-product-api/cmd/domain/entity"
)

//ProductService serves Product Repository func's
type ProductService struct {
	Repository *repository.ProductRepository
}

//Save creates a new product
func (s *ProductService) Save(p *entity.Product) int {
	return s.Repository.Save(p)
}
