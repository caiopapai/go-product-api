package repository

import (
	"database/sql"
	"fmt"

	"github.com/caiopapai/go-product-api/cmd/domain/entity"
)

//ProductRepository is a product repository
type ProductRepository struct {
	DB *sql.DB
}

//Save creates a new product
func (r *ProductRepository) Save(p *entity.Product) error {

	sqlStatement := `
	INSERT INTO pooble.product (name, price,due_date,brand_name,barcode,is_vegan, measurement_value, measurement_code)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	id := 0
	err := r.DB.QueryRow(sqlStatement, p.Name, p.Price, p.DueDate, p.BrandName,
		p.Barcode, p.IsVegan, p.Measurement.Code, p.Measurement.Value).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("New record ID is:", id)
	return nil
}

//GetByBarcode returns a Product filtered by Barcode
func GetByBarcode() entity.Product {
	return entity.Product{}
}

//GetByName returns a Product filtered by Name
func GetByName() entity.Product {
	return entity.Product{}
}

//Get return ALL Products
func Get() []entity.Product {
	return []entity.Product{}
}

//Remove one Product
func Remove(p *entity.Product) error {
	return nil
}

//Update one product
func Update(p *entity.Product) error {
	return nil
}
