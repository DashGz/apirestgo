package main

import (
	"github.com/gin-gonic/gin"
	"../controllers"
	"log"
)
const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main()  {

	router.GET("/users/:userId", controllers.GetUserFromAPI)
	router.GET("/countries/:countryId", controllers.GetCountryFromAPI)
	router.GET("/sites/:userId", controllers.GetSiteFromAPI)
	router.GET("/results/:userId", controllers.GetResultFromAPICh)
	log.Fatal(router.Run(port))
}
