package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
	"../utils"
)
const (
	paramUserID  = "userId"
)

func GetUserFromAPI(c *gin.Context)  {
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

	response, apiErr := services.GetUser(id)
	if err  != nil{
		c.JSON(apiErr.Status, err)
		return
	}

	c.JSON(http.StatusOK, response)
}