package repository

import (
	"alex/taxi-server/models"
	"github.com/guregu/dynamo"
	"time"
)

type IOrder interface {
	Create(order *models.InputOrder) error
	UpdateStatus(id, driverId, status string) error
	Get(id string) (*models.InputOrder, error)
}


type Order struct {
	db  *dynamo.DB
}

func (o *Order) Get(id string) (*models.InputOrder, error)  {
	table := o.db.Table("Order")
	var result models.InputOrder

	err := table.Get("Id", id).One(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (o *Order) UpdateStatus(id, driverId, status string) error {
	table := o.db.Table("Order")

	err := table.Update("Id", id).
		If("DriverId = ?", driverId).
		Set("Status", status).
		Set("LastUpdated", time.Now()).
		Run()

	return err
}

func (o *Order) Create(order *models.InputOrder) error {
	table := o.db.Table("Order")

	err := table.Put(order).Run()

	return err
}
