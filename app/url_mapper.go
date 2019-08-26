package app

import (
	"alex/taxi-server/handlers"
	"github.com/labstack/echo"
	"net/http"
)

type Mapper struct {
	//driver handlers
	DriverMapper *DriverMapper
}

func NewMapper(driverMapper *DriverMapper) *Mapper {
	return &Mapper{DriverMapper: driverMapper}
}

type DriverMapper struct {
	CreatorDriver echo.HandlerFunc
	GetterDriver  echo.HandlerFunc
	GetterItemDriver echo.HandlerFunc
}

func NewDriverMapper(driver *handlers.Driver) *DriverMapper {
	return &DriverMapper{
		CreatorDriver: driver.Create,
		GetterDriver: driver.GetAll,
		GetterItemDriver: driver.GetItem,
	}
}

func routingAPI(e *echo.Echo, mapper *Mapper){

	//ping
	e.GET("/ping", func(context echo.Context) error {
		return context.String(http.StatusOK, "pong")
	})


	driverGroup := e.Group("/drivers")
	driverGroup.POST("/", mapper.DriverMapper.CreatorDriver)
	driverGroup.GET("/", mapper.DriverMapper.GetterDriver)
	driverGroup.GET("/:driverId", mapper.DriverMapper.GetterItemDriver)
}
