package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sync"
)

type Site struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	CountryID          string   `json:"country_id"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	MercadopagoVersion int      `json:"mercadopago_version"`
	DefaultCurrencyID  string   `json:"default_currency_id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	Settings struct {
		IdentificationTypes      []interface{} `json:"identification_types"`
		TaxpayerTypes            []interface{} `json:"taxpayer_types"`
		IdentificationTypesRules []interface{} `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Categories []struct {
	} `json:"categories"`
}

func (site *Site) Get() *utils.ApiError {
	if site.ID == "" {
		return &utils.ApiError{
			"Site ID is empty.",
			http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s", utils.UrlSites, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {
		return &utils.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	return nil
}

func (site *Site) GetWG(waitgroup *sync.WaitGroup) *utils.ApiError {
	if site.ID == "" {
		return &utils.ApiError{
			"Site ID is empty.",
			http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s", utils.UrlSites, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {
		return &utils.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	waitgroup.Done()

	return nil
}

func (site *Site) GetCh(values chan Result) {
	if site.ID == "" {
		values <- Result{
			ApiError:&utils.ApiError{
				"Site ID is empty.",
				http.StatusBadRequest,
			},
		}
	}
	url := fmt.Sprintf("%s%s", utils.UrlSites, site.ID)

	response, err := http.Get(url)
	if err != nil {
		values <- Result{
			ApiError:&utils.ApiError{
				err.Error(),
				http.StatusInternalServerError,
			},
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		values <- Result{
			ApiError:&utils.ApiError{
				err.Error(),
				http.StatusInternalServerError,
			},
		}
	}

	if err := json.Unmarshal(data, &site); err != nil {
		values <- Result{
			ApiError:&utils.ApiError{
				err.Error(),
				http.StatusInternalServerError,
			},
		}
	}


	values <- Result{
		Site: site,
	}

}
