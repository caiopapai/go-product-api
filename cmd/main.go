package main

import "github.com/caiopapai/go-product-api/cmd/interface/web"

func main() {

	router := web.Router()
	router.Run(":8080")
}
