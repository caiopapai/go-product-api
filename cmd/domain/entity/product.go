package entity

//Product is our Domain Entity
type Product struct {
	ID          int
	Name        string      `json:"name"`
	Barcode     string      `json:"barcode"`
	BrandName   string      `json:"brandName"`
	IsVegan     bool        `json:"isVegan"`
	Price       string      `json:"price"`
	DueDate     string      `json:"dueDate"`
	Measurement Measurement `json:"measurement"`
}

//Measurement of product
type Measurement struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}
