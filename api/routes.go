package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"../classes"
)

func index(c *gin.Context) {
	c.Redirect(301, "/getAll/")
}

func getAll(c *gin.Context) {
	car1 := BaseClass.LightCar{}
	car2 := BaseClass.HeavyCar{}

	car1.Init("Toyota Starlet")
	car2.Init("Mazda Axela")

	c.JSON(http.StatusOK, []BaseClass.BaseInterface{&car1, &car2})
}
