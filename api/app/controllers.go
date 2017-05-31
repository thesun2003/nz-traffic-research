package app

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/thesun2003/traffic-simulation/api/db"
	"github.com/thesun2003/traffic-simulation/api/classes"
)

func index(c *gin.Context) {
	c.Redirect(301, "/getAll/")
}

func getAll(c *gin.Context) {
	// will be used later
	db.Config()

	car1 := BaseClass.LightCar{}
	car2 := BaseClass.HeavyCar{}

	car1.Init("Toyota Starlet")
	car2.Init("Mazda Axela")

	c.JSON(http.StatusOK, []BaseClass.BaseInterface{&car1, &car2})
}
