package main

import (
	"github.com/aliffadillah/mrt-schedule-golang.git/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	Initiaterouter()
}

func Initiaterouter() {
	var (
		router = gin.Default()
		api    = router.Group("v1/api")
	)

	station.Initiate(api)

	router.Run(":8080")
}
