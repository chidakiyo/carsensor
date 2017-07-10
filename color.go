package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Color Master

type ColorQuery struct {
	// apiキー
	Key string
}

type ColorResponse struct {
	Base
	Colors []Color `json:"color"`
}

type Color struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func SearchColor(param ColorQuery) ColorResponse {

	// TODO code利用する実装

	resp, err := http.Get(URL_COLOR + "?key=" + param.Key + "&format=" + FORMAT) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results ColorResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("[[ RESPONSE ]]\n%+v\n", response)

	return response.Results
}
