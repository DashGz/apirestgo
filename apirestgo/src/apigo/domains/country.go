package domains

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../utils"
	"sync"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	CurrencyID         string `json:"currency_id"`
	DecimalSeparator   string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
}

func (country *Country) Get() *utils.ApiError {
	if country.ID == "" {
		return &utils.ApiError{
			"Country ID is empty.",
			http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s", utils.UrlCountries, country.ID)

	response, err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil {
		return &utils.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	return nil
}

func (country *Country) GetWG(waitgroup *sync.WaitGroup) *utils.ApiError {
	if country.ID == "" {
		return &utils.ApiError{
			"Country ID is empty.",
			http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s", utils.UrlCountries, country.ID)

	response, err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil {
		return &utils.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	waitgroup.Done()

	return nil
}

func (country *Country) GetCh(values chan Result)  {
	if country.ID == "" {

	}
	url := fmt.Sprintf("%s%s", utils.UrlCountries, country.ID)

	response, err := http.Get(url)
	if err != nil {

	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {

	}

	if err := json.Unmarshal(data, &country); err != nil {

	}

	values <- Result{
		Country: country,
	}
}