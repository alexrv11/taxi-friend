package handlers

import (
	"alex/taxi-server/models"
	"alex/taxi-server/services"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Driver struct {
	DriverService services.IDriver
}

func NewDriver(driverService services.IDriver) *Driver {
	return &Driver{DriverService:driverService}
}

func (d *Driver) Create(c echo.Context) error {

	input := new(models.InputDriver)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	d.DriverService.Create(input)
	return c.JSON(http.StatusCreated, nil)
}

func (d *Driver) GetAll(c echo.Context) error {

	radio, _ := strconv.ParseFloat(c.QueryParam("radio"), 32)
	latitude, _ := strconv.ParseFloat(c.QueryParam("latitude"), 32)
	longitude, _ := strconv.ParseFloat(c.QueryParam("longitude"), 32)
	result, err := d.DriverService.GetAll(radio, latitude, longitude)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, models.Response{Result: result})
}

func (d *Driver) GetItem(c echo.Context) error {
	driverId := c.Param("driverId")
	result, err := d.DriverService.GetItem(driverId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, models.Response{Result: result})
}
