package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestCatalog(t *testing.T) {

	fmt.Println("---> Catalog -----------------------------")

	param := CatalogQuery{
		Key: API_KEY,
	}

	actual := SearchCatalog(param)

	fmt.Println(actual)

	fmt.Println("==========")
}
