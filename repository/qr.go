package repository

import (
	"alex/taxi-server/models"
	"github.com/guregu/dynamo"
)

type IQr interface {
	Get(code string) (*models.Qr, error)
	UpdateDriver(driverID, qr string)  error
	Create(qr models.Qr) error
}

type Qr struct {
	db  *dynamo.DB
}

func (q *Qr) Get(code string) (*models.Qr, error)  {
	table := q.db.Table("CreditQr")
	var result models.Qr

	err := table.Get("Code", code).One(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (q *Qr) UpdateDriver(driverID, qr string) error {
	table := q.db.Table("CreditQr")

	err := table.Update("Code", qr).
		If("DriverID = ?", "none").
		Set("DriverID", driverID).
		Run()

	return err
}

func (q *Qr) Create(qr models.Qr) error {
	table := q.db.Table("CreditQr")

	err := table.Put(qr).Run()

	return err
}