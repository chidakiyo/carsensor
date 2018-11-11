package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestCatalog(t *testing.T) {

	fmt.Println("---> Catalog -----------------------------")

	param := CatalogQuery{
		Key: API_KEY,
		Country:"GER",
	}

	actual, _ := SearchCatalog(param)

	fmt.Println(PrettyPrint(actual))

	fmt.Println("==========")
}
