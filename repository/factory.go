package repository

import "github.com/guregu/dynamo"

type IFactory interface {
	CreateDriverRepository() IDriver
}

type Factory struct {
	DB *dynamo.DB
}

func (f *Factory) CreateDriverRepository() IDriver {
	return &Driver{f.DB}
}
