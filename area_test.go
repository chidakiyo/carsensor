package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestSearchArea(t *testing.T) {

	fmt.Println("---> Area Master -----------------------------")

	param := AreaQuery{
		Key:  API_KEY,
		Code: "",
	}

	actual := SearchArea(param)
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(actual)

	fmt.Println("==========")
}
