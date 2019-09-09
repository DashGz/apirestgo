package services

import (
	"../domains"
	"../utils"
	"sync"
)

func GetResult(userId int) (*domains.Result, *utils.ApiError) {

	user := domains.User{
		ID: userId,
	}

	if err := user.Get(); err != nil{
		return nil, err
	}

	country := domains.Country{
		ID: user.CountryID,
	}

	site := domains.Site{
		ID: user.SiteID,
	}

	if err := country.Get(); err != nil{
		return nil, err
	}

	if err := site.Get(); err != nil{
		return nil, err
	}
	//country.Get()
	//site.Get()

	response := domains.Result{
		User: &user,
		Site: &site,
		Country: &country,
	}

	return &response, nil
}


func GetResultWG(userId int) (*domains.Result, *utils.ApiError) {

	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	user := domains.User{
		ID: userId,
	}

	if err := user.Get(); err != nil{
		return nil, err
	}

	country := domains.Country{
		ID: user.CountryID,
	}

	site := domains.Site{
		ID: user.SiteID,
	}


	go country.GetWG(&waitgroup)
	go site.GetWG(&waitgroup)

	waitgroup.Wait()

	response := domains.Result{
		User: &user,
		Site: &site,
		Country: &country,
	}

	return &response, nil
}

func GetResultCh(userId int) (*domains.Result, *utils.ApiError) {

	user := domains.User{
		ID: userId,
	}

	if err := user.Get(); err != nil{
		return nil, err
	}

	country := domains.Country{
		ID: user.CountryID,
	}

	site := domains.Site{
		ID: user.SiteID,
	}

	values := make(chan domains.Result, 2)
	defer close(values)


	go country.GetCh(values)
	go site.GetCh(values)

	response := domains.Result{
		User: &user,
	}

	for i := 0; i < 2; i++{
		result := <- values
		if result.Site != nil{
			response.Site = result.Site
		}
		if result.Country != nil{
			response.Country = result.Country
		}
	}



	return &response, nil
}