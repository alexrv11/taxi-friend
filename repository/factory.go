package repository

import "github.com/guregu/dynamo"

type IFactory interface {
	CreateDriverRepository() IDriver
	CreateQrRepository() IQr
	CreateOrderRepository() IOrder
}

type Factory struct {
	DB *dynamo.DB
}

func (f *Factory) CreateDriverRepository() IDriver {
	return &Driver{f.DB }
}

func (f *Factory) CreateQrRepository() IQr {
	return &Qr{f.DB }
}

func (f *Factory) CreateOrderRepository() IOrder {
	return &Order{f.DB }
}
