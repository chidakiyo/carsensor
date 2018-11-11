package carsensor_api_go

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// Color Master

type ColorQuery struct {
	// apiキー
	Key string `json:"key"`
}

type ColorResponse struct {
	Base
	Colors []Color `json:"color"`
}

type Color struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func SearchColor(param ColorQuery) (ColorResponse, error) {

	u, err := CreateURL(URL_COLOR, param)
	if err != nil {
		err := errors.Wrap(err, "create url fail.")
		return ColorResponse{}, err // fail
	}

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		err := errors.Wrap(err, "http get fail.")
		return ColorResponse{}, err // fail
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results ColorResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return ColorResponse{}, err // fail
	}

	return response.Results, nil // success
}
