package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Bland Master

type BodyTypeQuery struct {
	// apiキー
	Key string
}

type BodyTypeResponse struct {
	Base
	BodyType []BodyType `json:"body"`
}

type BodyType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func SearchBodyType(param BodyTypeQuery) BodyTypeResponse {

	// TODO code利用する実装

	resp, err := http.Get(URL_BODY_TYPE + "?key=" + param.Key + "&format=" + FORMAT) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results BodyTypeResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("[[ RESPONSE ]]\n%+v\n", response)

	return response.Results
}
