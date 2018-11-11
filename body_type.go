package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// Bland Master

type BodyTypeQuery struct {
	// apiキー
	Key string `json:"key"`
}

type BodyTypeResponse struct {
	Base
	BodyType []BodyType `json:"body"`
}

type BodyType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func SearchBodyType(param BodyTypeQuery) (BodyTypeResponse, error) {

	u, err := CreateURL(URL_BODY_TYPE, param)
	if err != nil {
		err := errors.Wrap(err, "create url fail.")
		return BodyTypeResponse{}, err // fail
	}

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		err := errors.Wrap(err, "http get fail.")
		return BodyTypeResponse{}, err // fail
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results BodyTypeResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return BodyTypeResponse{}, err // fail
	}

	fmt.Printf("[[ RESPONSE ]]\n%+v\n", response)

	return response.Results, nil // success
}
