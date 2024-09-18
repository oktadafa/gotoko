package app

import "github.com/oktadafa/gotoko/app/models"

type Models struct{
	Model interface{}
}

func RegiseterModels() []Models{
	return []Models {
		{Model:models.User{}},
		{Model:models.Address{}},
		{Model:models.Category{}},
		{Model:models.Order{}},
		{Model:models.OrderCustomer{}},
		{Model:models.OrderItem{}},
		{Model:models.Payment{}},
		{Model:models.Product{}},
		{Model:models.ProductImage{}},
		{Model:models.Section{}},
		{Model:models.Shipment{}},
	}
}

