package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
	"../utils"
	"../utils/circuit_breaker"
)


func GetResultFromAPI(c *gin.Context)  {
	userID := c.Param(paramUserID)

	id, err := strconv.Atoi(userID)
	if err  != nil{
		apiErr := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, err)
		return
	}

	response, apiErr := services.GetResult(id)
	if err  != nil{
		c.JSON(apiErr.Status, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetResultFromAPIWG(c *gin.Context)  {
	userID := c.Param(paramUserID)

	id, err := strconv.Atoi(userID)
	if err  != nil{
		apiErr := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, err)
		return
	}

	response, apiErr := services.GetResultWG(id)
	if err  != nil{
		c.JSON(apiErr.Status, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetResultFromAPICh(c *gin.Context)  {
	userID := c.Param(paramUserID)

	cb := new(circuit_breaker.CircuitBreaker)
	cb.StartCB("CLOSED", 5,3)


	id, err := strconv.Atoi(userID)
	if err  != nil{
		apiErr := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, err)
		return
	}


	if cb.State == "CLOSED" {
		response, apiErr := services.GetResultCh(id)
		if err != nil {
			cb.RecordFailure()
			c.JSON(apiErr.Status, err)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
