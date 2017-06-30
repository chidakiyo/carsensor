package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Country Master

const (
	FORMAT = "json"
)

type CountryRequest struct {
	Key  string
	Code string
}

type CountryResponse struct {
	ApiVersion       string    `json:"api_version"`
	ResultsReturned  string    `json:"results_returned"`
	ResultsAvailable int64     `json:"results_available"`
	ResultsStart     int64     `json:"results_start"`
	Countries        []Country `json:"country"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Root struct {
	Results CountryResponse `json:"results"`
}

func SearchCountry(param CountryRequest) string {
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
