package carsensor_api_go

import (
	"fmt"
	"os"
	"testing"
)

var API_KEY = os.Getenv("CS_KEY") // 環境変数にKEYを設定しておく ex. export CS_KEY={YOUR_KEY}

func TestExecute(t *testing.T) {

	param := CountryRequest{
		Key:  API_KEY,
		Code: "",
	}

	actual := Execute(param)
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(actual)
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
