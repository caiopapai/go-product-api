package product

//Data is the struct that holds the Object
type Data struct {
	Product Product `json:"product"`
}

//Product is our Domain Entity
type Product struct {
	Name        string      `json:"name"`
	Barcode     string      `json:"barcode"`
	BandName    string      `json:"bandName"`
	IsVegan     string      `json:"isVegan"`
	Measurement Measurement `json:"measurement"`
}

//Measurement is the unit of mesurement
type Measurement struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}
