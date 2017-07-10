package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Area Master

type AreaQuery struct {
	// apiキー
	Key string
	// ブランドコード
	Code string
}

type AreaResponse struct {
	Base
	Areas []Area `json:"large_area"`
}

type Area struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func SearchArea(param AreaQuery) AreaResponse {

	// TODO code利用する実装

	resp, err := http.Get(URL_AREA + "?key=" + param.Key + "&format=" + FORMAT) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results AreaResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("[[ RESPONSE ]]\n%+v\n", response)

	return response.Results
}
