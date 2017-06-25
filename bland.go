package carsensor_api_go

// Bland Master

const (
//URL = "http://webservice.recruit.co.jp/carsensor/brand/v1/"
)

type BlandRequest struct {
	// apiキー
	Key string
	// ブランドコード
	Code string
	// 国コード
	Country string
}

type Base struct {
	ApiVersion       string `json:"api_version"`
	ResultsAvailable string `json:"results_available"`
	ResultsReturned  string `json:"results_returned"`
	ResultsStart     string `json:"results_start"`
}

type BlandResponse struct {
	Base
	Brands []Bland
}

type Bland struct {
	code string
	name string
	Country
}
