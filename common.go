package carsensor_api_go

type Base struct {
	ApiVersion       string `json:"api_version"`
	ResultsAvailable int64  `json:"results_available"`
	ResultsReturned  string `json:"results_returned"`
	ResultsStart     int64  `json:"results_start"`
}
