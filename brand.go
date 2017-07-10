package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Bland Master

type BrandQuery struct {
	// apiキー
	Key string
	// ブランドコード
	Code string
	// 国コード
	Country string
}

type BrandResponse struct {
	Base
	Brands []Brand `json:"brand"`
}

type Brand struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Country
}

func SearchBrand(param BrandQuery) BrandResponse {

	// TODO code利用する実装

	resp, err := http.Get(URL_BRAND + "?key=" + param.Key + "&format=" + FORMAT) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results BrandResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("[[ RESPONSE ]]\n%+v\n", response)

	return response.Results
}
