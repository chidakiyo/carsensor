package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// catalog api

type CatalogQuery struct {
	//
	Key string
	//
	Brand Brand
	//
	Model string
	//
	Country Country
	//
	Body BodyType
	//
	Person int64
	//
	YearOld int64
	//
	YearNew int64
	//
	Welfare string
	//
	Series string
	//
	Keyword string
	//
	WidthMax string
	//
	HeightMax string
	//
	LengthMax string
	//
	Order string
	//
	Start string
	//
	Count string
	//
	Format string
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

func SearchCatalog(param CatalogQuery) CatalogResponse {

	// TODO code利用する実装

	resp, err := http.Get(URL_CATALOG + "?key=" + param.Key + "&keyword=インプレッサ" + "&format=" + FORMAT) // TODO 全件
	defer resp.Body.Close()
	if err != nil {
		// TODO handle error
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s", body)

	response := &struct {
		Results CatalogResponse `json:"results"`
	}{}
	if err := json.Unmarshal(body, response); err != nil {
		panic(err)
	}

	fmt.Printf("[[ RESPONSE ]]\n%+v\n", response)

	return response.Results
}
