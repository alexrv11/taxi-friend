package handlers

import (
	"alex/taxi-server/models"
	"alex/taxi-server/services"
	"github.com/labstack/echo"
	"net/http"
)

type Qr struct {
	QrService services.IQr
}

func (q *Qr) Create(c echo.Context) error {

	input := new(models.Qr)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	q.QrService.Create(*input)
	return c.JSON(http.StatusCreated, nil)
}

func (q *Qr) Get(c echo.Context) error {
	qrCode := c.Param("code")

	result, _ := q.QrService.GetItem(qrCode)

	if result == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, result)
}

func (q *Qr) UpdateDriver(c echo.Context) error {
	qrCode := c.Param("code")
	driverId := c.Param("driverId")

	err := q.QrService.UpdateDriver(driverId, qrCode)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
