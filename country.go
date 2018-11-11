package carsensor_api_go

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// Country Master

type CountryQuery struct {
	Key  string `json:"key"`
	Code string `json:"code"`
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

func SearchCountry(param CountryQuery) (CountryResponse, error) {

	u, err := CreateURL(URL_COUNTRY, param)
	if err != nil {
		err := errors.Wrap(err, "create url fail.")
		return CountryResponse{}, err // fail
	}

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		err := errors.Wrap(err, "http get fail.")
		return CountryResponse{}, err // fail
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results CountryResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return CountryResponse{}, err // fail
	}

	return response.Results, nil // success
}
