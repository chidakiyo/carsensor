package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestSearchBrand(t *testing.T) {

	param := BlandRequest{
		Key:  API_KEY,
		Code: "",
	}

	actual := SearchBrand(param)
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(actual)
}
