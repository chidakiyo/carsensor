package carsensor_api_go

import (
	"fmt"
	"os"
	"testing"
)

var API_KEY = os.Getenv("CS_KEY") // 環境変数にKEYを設定しておく ex. export CS_KEY={YOUR_KEY}

func TestSearchCountry(t *testing.T) {

	fmt.Println("---> Country Master -----------------------------")

	param := CountryQuery{
		Key:  API_KEY,
		Code: "",
	}

	actual, _ := SearchCountry(param)
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
