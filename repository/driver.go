package repository

import (
	"alex/taxi-server/models"
	"github.com/guregu/dynamo"
)

type IDriver interface {
	Create(driver *models.InputDriver) error
	GetAll() ([]models.Driver, error)
	GetItem(driverId string) (*models.Driver, error)
}

type Driver struct {
	db  *dynamo.DB
}

func (d *Driver) Create(driver *models.InputDriver) error {
	table := d.db.Table("Driver")

	err := table.Put(driver).Run()

	return err
}

func (d *Driver) GetAll() ([]models.Driver, error) {
	table := d.db.Table("Driver")

	// get all items
	var results []models.Driver
	err := table.Scan().All(&results)

	return results, err
}

func (d *Driver) GetItem(driverId string) (*models.Driver, error) {
	table := d.db.Table("Driver")
	var result models.Driver

	err := table.Get("Id", driverId).One(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
