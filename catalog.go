package carsensor_api_go

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// catalog api

type CatalogQuery struct {
	//
	Key string `json:"key"`
	//
	Brand string `json:"brand"`
	//
	Model string `json:"model"`
	//
	Country string `json:"country"`
	//
	Body string `json:"body"`
	//
	Person int64 `json:"person"`
	//
	YearOld int64 `json:"year_old"`
	//
	YearNew int64 `json:"year_new"`
	//
	Welfare string `json:"welfare"`
	//
	Series string `json:"series"`
	//
	Keyword string `json:"keyword"`
	//
	WidthMax string `json:"width_max"`
	//
	HeightMax string `json:"height_max"`
	//
	LengthMax string `json:"length_max"`
	//
	Order string `json:"order"`
	//
	Start string `json:"start"`
	//
	Count string `json:"count"`
	//
	Format string `json:"format"`
}

type CatalogResponse struct {
	Base
	Catalogs []Catalog `json:"catalog"`
}

type Catalog struct {
	Brand       Brand    `json:"brand"`
	Model       string   `json:"model"`
	Grade       string   `json:"grade"`
	Price       int64    `json:"price"`
	Description string   `json:"desc"`
	Body        BodyType `json:"body"`
	Person      string   `json:"person"`
	Period      string   `json:"period"`
	Series      string   `json:"series"`
	Width       string   `json:"width"`
	Height      string   `json:"height"`
	Length      string   `json:"length"`
	Photo       PhotoAll `json:"photo"`
	URLS        URLS     `json:"urls"`
	Desc        string   `json:"desc"`
}

func SearchCatalog(param CatalogQuery) (CatalogResponse, error) {

	u, err := CreateURL(URL_CATALOG, param)
	if err != nil {
		err := errors.Wrap(err, "create url fail.")
		return CatalogResponse{}, err // fail
	}

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		err := errors.Wrap(err, "http get fail.")
		return CatalogResponse{}, err // fail
	}

	body, err := ioutil.ReadAll(resp.Body)

	response := &struct {
		Results CatalogResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return CatalogResponse{}, err // fail
	}

	return response.Results, nil // success
}
