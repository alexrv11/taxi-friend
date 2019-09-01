package repository

import "github.com/guregu/dynamo"

type IFactory interface {
	CreateDriverRepository() IDriver
	CreateQrRepository() IQr
}

type Factory struct {
	DB *dynamo.DB
}

func (f *Factory) CreateDriverRepository() IDriver {
	return &Driver{f.DB}
}

func (f *Factory) CreateQrRepository() IQr {
	return &Qr{f.DB}
}
