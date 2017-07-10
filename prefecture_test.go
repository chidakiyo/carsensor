package carsensor_api_go

import (
	"fmt"
	"testing"
)

func TestSearchPrefecture(t *testing.T) {

	fmt.Println("---> Prefecture Master -----------------------------")

	param := PrefectureQuery{
		Key:  API_KEY,
		Code: "",
	}

	actual := SearchPrefecture(param)
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(actual)

	fmt.Println("==========")
}
