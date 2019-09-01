package services

import (
	"alex/taxi-server/models"
	"alex/taxi-server/repository"
	"alex/taxi-server/utils"
	"github.com/google/uuid"
)

type Driver struct {
	driverRepository repository.IDriver
}

func NewDriver(repoFactory repository.IFactory) *Driver {
	return &Driver{driverRepository: repoFactory.CreateDriverRepository()}
}

func (d *Driver) Create(driver *models.InputDriver) error  {
	driver.Password = utils.CreatePassword(driver.Password)
	driver.Id = uuid.New().String()
	driver.Status = "Registered"
	driver.Credit = 0
	return d.driverRepository.Create(driver)
}

func (d *Driver) GetAll(radio, latitude, longitude float64) ([]models.DriverLocation, error) {
	drivers, err := d.driverRepository.GetAll()

	if err != nil {
		return nil, err
	}

	result := make([]models.DriverLocation, 0)
	for _, driver := range drivers {
		if radio >= utils.DistanceInKmBetweenEarthCoordinates(latitude, longitude, driver.Latitude, driver.Longitude) {
			item := models.DriverLocation{
				ID:driver.ID,
				Latitude: driver.Latitude,
				Longitude: driver.Longitude,
				Status: driver.Status,
				Direction: driver.Direction,
			}
			result = append(result, item)
		}
	}

	return result, nil
}

func (d *Driver) GetItem(driverId string) (*models.Driver, error)  {
	return d.driverRepository.GetItem(driverId)
}

func (d *Driver) UpdateLocation(driverId string, location models.Location) error {

	return d.driverRepository.UpdateLocation(driverId, location)
}