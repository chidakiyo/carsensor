package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestUsedCar(t *testing.T) {

	fmt.Println("---> UsedCar Search -----------------------------")

	param := UsedCarQuery{
		Key:     API_KEY,
		Mission: 1, // AT
		Model:   "ジュリア",
		Count:   1,
	}

	actual, err := SearchUsedCar(param)
	if err != nil {
		t.Fatalf("%s", err)
	}
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(PrettyPrint(actual))

	fmt.Println("==========")
}

//func TestExecute2(t *testing.T) {
//
//	r := Root{
//		Results: Results{
//			ApiVersion: "B",
//		},
//	}
//
//	b, _ := json.Marshal(r)
//
//	fmt.Println(string(b) + "----")
//
//}
