package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestSearchColor(t *testing.T) {

	fmt.Println("---> Color Master -----------------------------")

	param := ColorQuery{
		Key: API_KEY,
	}

	actual, _ := SearchColor(param)
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(PrettyPrint(actual))

	fmt.Println("==========")
}
