package carsensor_api_go

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// Area Master

type AreaQuery struct {
	// apiキー
	Key string `json:"key"`
	// ブランドコード
	Code string `json:"code"`
}

type AreaResponse struct {
	Base
	Areas []Area `json:"large_area"`
}

type Area struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func SearchArea(param AreaQuery) (AreaResponse, error) {

	u, err := CreateURL(URL_AREA, param)
	if err != nil {
		err := errors.Wrap(err, "create url fail.")
		return AreaResponse{}, err // fail
	}

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		err := errors.Wrap(err, "http get fail.")
		return AreaResponse{}, err // fail
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results AreaResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return AreaResponse{}, err // fail
	}


	return response.Results, nil // success
}
