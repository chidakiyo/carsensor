package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestSearchBrand(t *testing.T) {

	fmt.Println("---> Brand Master -----------------------------")

	param := BrandQuery{
		Key:  API_KEY,
		Code: "",
	}

	actual, _ := SearchBrand(param)
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(PrettyPrint(actual))

	fmt.Println("==========")
}
