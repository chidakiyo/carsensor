package carsensor_api_go

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

// Bland Master

const (
//URL = "http://webservice.recruit.co.jp/carsensor/brand/v1/"
)

type BlandRequest struct {
	// apiキー
	Key string
	// ブランドコード
	Code string
	// 国コード
	Country string
}

type Base struct {
	ApiVersion       string `json:"api_version"`
	ResultsAvailable string `json:"results_available"`
	ResultsReturned  string `json:"results_returned"`
	ResultsStart     string `json:"results_start"`
}

type BlandResponse struct {
	Base
	Brands []Bland
}

type Bland struct {
	code string
	name string
	Country
}

func SearchBrand(param BlandRequest) string {
	URL := "http://webservice.recruit.co.jp/carsensor/country/v1/"

	// TODO code利用する実装

	resp, err := http.Get(URL + "?key=" + param.Key + "&format=" + FORMAT) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &Root{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", response)

	return string(body)
}
