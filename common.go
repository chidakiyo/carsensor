package carsensor_api_go

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"path"
)

type Base struct {
	ApiVersion       string `json:"api_version"`
	ResultsAvailable int64  `json:"results_available"`
	ResultsReturned  string `json:"results_returned"`
	ResultsStart     int64  `json:"results_start"`
}

func QueryToMap(query interface{}) (map[string]interface{}, error) {
	paramMap := map[string]interface{}{}

	tmpJson, err := json.Marshal(query)
	if err != nil {
		err := errors.Wrap(err, "json marshal fail.")
		return paramMap, err // fail
	}

	if err := json.Unmarshal(tmpJson, &paramMap); err != nil {
		err := errors.Wrap(err, "json unmarshal fail.")
		return paramMap, err // fail
	}

	return paramMap, nil // success
}

func CreateURL(p string, query interface{}) (*url.URL, error) {

	u, err := url.Parse(CARSENSOR)
	if err != nil {
		err := errors.Wrap(err, "url parse fail.")
		return u, err
	}

	u.Path = path.Join(u.Path, p)

	mappedQuery, err := QueryToMap(query)
	if err != nil {
		err := errors.Wrap(err, "query to map fail.")
		return u, err
	}

	q := u.Query()
	for k, v := range mappedQuery {
		switch x := v.(type) {
		case int64:
			if v.(int64) > 0 {
				q.Add(k, fmt.Sprintf("%v", v)) // TODO
			}
		case float64:
			if v.(float64) > 0 {
				q.Add(k, fmt.Sprintf("%v", v)) // TODO
			}
		case string:
			if len(v.(string)) > 0 {
				q.Add(k, v.(string))
			}
		default:
			panic(fmt.Sprintf("unknown type. fixme %#+v", x))
		}
	}
	q.Add("format", FORMAT)
	u.RawQuery = q.Encode()

	fmt.Printf("URL : %s", u.String()) // TODO debug

	return u, nil
}

func PrettyPrint(object interface{}) string {

	jsonString, _ := json.MarshalIndent(object, "", "  ")

	return string(jsonString)

}
