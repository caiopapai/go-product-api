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

//Save creates a new product and return ID
func (r *ProductRepository) Save(p *entity.Product) int {

	sqlStatement := ` INSERT INTO pooble.product (name, 
												  price,
												  due_date,
												  brand_name,
												  barcode,
												  is_vegan, 
												  measurement_value, 
												  measurement_code)
					  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
					  RETURNING Id`

	id := 0

	err := r.DB.QueryRow(sqlStatement, p.Name, p.Price, p.DueDate, p.BrandName,
		p.Barcode, p.IsVegan, p.Measurement.Value, p.Measurement.Code).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	fmt.Println("New record ID is:", &id)
	return id
}

//GetByID returns a Product filtered by ID
func GetByID(id int) entity.Product {
	return entity.Product{}
}

//GetByBarcode returns a Product filtered by Barcode
func GetByBarcode() entity.Product {
	return entity.Product{}
}

//GetByName returns a Product filtered by Name
func GetByName() entity.Product {
	return entity.Product{}
}

//GetAll return ALL Products
func (r *ProductRepository) GetAll() []entity.Product {

	rows, err := r.DB.Query("SELECT name, price, due_date, brand_name, barcode, is_vegan, measurement_value, measurement_code FROM pooble.product")
	if err != nil {
		panic(err.Error())
	}

	products := []entity.Product{}

	for rows.Next() {
		var name string
		var barcode string
		var brandname string
		var isvegan bool
		var price string
		var duedate string
		var mescode string
		var mesvalue string

		err = rows.Scan(&name, &price, &duedate, &brandname, &barcode, &isvegan, &mesvalue, &mescode)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		p := entity.Product{
			Name:        name,
			Barcode:     barcode,
			BrandName:   brandname,
			DueDate:     duedate,
			IsVegan:     isvegan,
			Price:       price,
			Measurement: entity.Measurement{Code: mescode, Value: mesvalue},
		}

		products = append(products, p)
	}

	return products
}

//Remove one Product
func Remove(p *entity.Product) error {
	return nil
}

//Update one product
func Update(p *entity.Product) error {
	return nil
}
