package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Country Master

type CountryQuery struct {
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

func SearchCountry(param CountryQuery) CountryResponse {

	// TODO code利用する実装

	resp, err := http.Get(URL_COUNTRY + "?key=" + param.Key + "&format=" + FORMAT) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results CountryResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("[[ RESPONSE ]]\n%+v\n", response)

	return response.Results
}
