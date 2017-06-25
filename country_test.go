package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {

	param := CountryRequest{
		Key:  "YOUR_KEY",
		Code: "",
	}

	actual := Execute(param)
	//if actual != expected {
	//	t.Errorf("got %v\nwant %v", actual, expected)
	//}

	fmt.Println(actual)
}

func TestExecute2(t *testing.T) {

	r := Root{
		Results: Results{
			ApiVersion: "B",
		},
	}

	b, _ := json.Marshal(r)

	fmt.Println(string(b) + "----")

}
