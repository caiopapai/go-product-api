package entity

//Product is our Domain Entity
type Product struct {
	Name        string      `json:"name"`
	Barcode     string      `json:"barcode"`
	BrandName   string      `json:"brandName"`
	IsVegan     bool        `json:"isVegan"`
	Price       string      `json:"price"`
	DueDate     string      `json:"dueDate"`
	Measurement measurement `json:"measurement"`
}

type measurement struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}
