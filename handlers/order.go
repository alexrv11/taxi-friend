package handlers

import (
	"alex/taxi-server/models"
	"alex/taxi-server/services"
	"github.com/labstack/echo"
	"net/http"
)

type Order struct {
	OrderService services.IOrder
}

func (o *Order) Create(c echo.Context) error {

	input := new(models.InputOrder)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	o.OrderService.Create(input)
	return c.JSON(http.StatusCreated, nil)
}

func (o *Order) Get(c echo.Context) error {
	id := c.Param("id")

	result, _ := o.OrderService.GetItem(id)

	if result == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, result)
}

func (o *Order) Update(c echo.Context) error {
	id := c.Param("id")

	input := new(models.InputOrder)
	input.ID = id
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := o.OrderService.Update(input)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}