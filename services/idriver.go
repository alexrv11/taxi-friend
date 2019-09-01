package services

import "alex/taxi-server/models"

type IDriver interface {
	Create(driver *models.InputDriver) error
	GetAll(radio, latitude, longitude float64) ([]models.DriverLocation,error)
	GetItem(driverId string) (*models.Driver, error)
	UpdateLocation(driverId string, location models.Location) error
}
