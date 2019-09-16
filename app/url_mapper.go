package app

import (
	"alex/taxi-server/handlers"
	"github.com/labstack/echo"
	"net/http"
)

type Mapper struct {
	//driver handlers
	Driver *DriverMapper
	Qr     *QrMapper
	Order *OrderMapper
}

func NewMapper(driverMapper *DriverMapper, qr *QrMapper, order *OrderMapper) *Mapper {
	return &Mapper{Driver: driverMapper, Qr: qr, Order: order}
}

type DriverMapper struct {
	CreatorDriver echo.HandlerFunc
	GetterDriver  echo.HandlerFunc
	GetterItemDriver echo.HandlerFunc
	UpdateLocation echo.HandlerFunc
}

type QrMapper struct {
	Create echo.HandlerFunc
	Get echo.HandlerFunc
	UpdateDriver echo.HandlerFunc
}

type OrderMapper struct {
	Create echo.HandlerFunc
	Get echo.HandlerFunc
	Update echo.HandlerFunc
}

func NewDriverMapper(driver *handlers.Driver) *DriverMapper {
	return &DriverMapper{
		CreatorDriver: driver.Create,
		GetterDriver: driver.GetAll,
		GetterItemDriver: driver.GetItem,
		UpdateLocation: driver.UpdateLocation,
	}
}

func NewQrMapper(qr *handlers.Qr) *QrMapper  {
	return &QrMapper{ Create: qr.Create, Get: qr.Get, UpdateDriver: qr.UpdateDriver }
}

func NewOrderMapper(order *handlers.Order) *OrderMapper {
	return &OrderMapper{ Create: order.Create, Get: order.Get, Update: order.Update }
}

func routingAPI(e *echo.Echo, mapper *Mapper){

	//ping
	e.GET("/ping", func(context echo.Context) error {
		return context.String(http.StatusOK, "pong")
	})


	driverGroup := e.Group("/drivers")
	driverGroup.POST("/", mapper.Driver.CreatorDriver)
	driverGroup.GET("/", mapper.Driver.GetterDriver)
	driverGroup.GET("", mapper.Driver.GetterDriver)
	driverGroup.GET("/:driverId", mapper.Driver.GetterItemDriver)
	driverGroup.PUT("/:driverId/location", mapper.Driver.UpdateLocation)

	qrGroup := e.Group("/qr")
	qrGroup.POST("/", mapper.Qr.Create)
	qrGroup.GET("/:code", mapper.Qr.Get)
	qrGroup.PUT("/:code/driver/:driverId", mapper.Qr.UpdateDriver)

	orderGroup := e.Group("/orders")
	orderGroup.POST("/", mapper.Order.Create)
	orderGroup.GET("/:id", mapper.Order.Get)
	orderGroup.PUT("/:id", mapper.Order.Update)
}
