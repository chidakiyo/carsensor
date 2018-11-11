package carsensor_api_go

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// Bland Master

type BrandQuery struct {
	// apiキー
	Key string `json:"key"`
	// ブランドコード
	Code string `json:"code"`
	// 国コード
	Country string `json:"country"`
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

func SearchBrand(param BrandQuery) (BrandResponse, error) {

	u, err := CreateURL(URL_BRAND, param)
	if err != nil {
		err := errors.Wrap(err, "create url fail.")
		return BrandResponse{}, err // fail
	}

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		err := errors.Wrap(err, "http get fail.")
		return BrandResponse{}, err // fail
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results BrandResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return BrandResponse{}, err // fail
	}

	return response.Results, nil // success
}
