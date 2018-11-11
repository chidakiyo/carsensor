package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestBodyType(t *testing.T) {

	fmt.Println("---> Body Type -----------------------------")

	param := BodyTypeQuery{
		Key: API_KEY,
	}

	actual, _ := SearchBodyType(param)
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(PrettyPrint(actual))

	fmt.Println("==========")
}
