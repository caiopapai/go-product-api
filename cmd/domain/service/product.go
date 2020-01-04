package service

import (
	"github.com/caiopapai/go-product-api/cmd/domain/repository"
	"github.com/caiopapai/go-product-api/cmd/infra/connection"

	"github.com/caiopapai/go-product-api/cmd/domain/entity"
)

//ProductService serves Product Repository func's
type ProductService struct {
	Repository *repository.ProductRepository
}

//Init initialize a new Product Service with a new DB connection
func Init() *ProductService {
	repo := repository.ProductRepository{DB: connection.GetConnection()}
	service := ProductService{Repository: &repo}
	return &service
}

//Save creates a new product
func (s *ProductService) Save(p *entity.Product) int {
	return s.Repository.Save(p)
}

//GetAll get product
func (s *ProductService) GetAll() []entity.Product {
	return s.Repository.GetAll()
}

//Remove removes one product
func (s *ProductService) Remove(id int) {
	productID := id
	s.Repository.Remove(productID)
}
