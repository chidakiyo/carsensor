package carsensor_api_go

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// Prefecture Master

type PrefectureQuery struct {
	// apiキー
	Key string `json:"key"`
	// ブランドコード
	Code string `json:"code"`
	// 大エリア
	Area string `json:"area"`
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

func SearchPrefecture(param PrefectureQuery) (PrefectureResponse, error) {

	u, err := CreateURL(URL_PREFECTURE, param)
	if err != nil {
		err := errors.Wrap(err, "create url fail.")
		return PrefectureResponse{}, err // fail
	}

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		err := errors.Wrap(err, "http get fail.")
		return PrefectureResponse{}, err // fail
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results PrefectureResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return PrefectureResponse{}, err // fail
	}

	return response.Results, nil // success
}
