package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Prefecture Master

type PrefectureQuery struct {
	// apiキー
	Key string
	// ブランドコード
	Code string
	// 大エリア
	Area string
}

type PrefectureResponse struct {
	Base
	Prefectures []Prefecture `json:"pref"`
}

type Prefecture struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Area Area   `json:"large_area"`
}

func SearchPrefecture(param PrefectureQuery) PrefectureResponse {

	// TODO code利用する実装

	// TODO large_area

	resp, err := http.Get(URL_PREFECTURE + "?key=" + param.Key + "&format=" + FORMAT) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results PrefectureResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("[[ RESPONSE ]]\n%+v\n", response)

	return response.Results
}
