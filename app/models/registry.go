package models

type Models struct {
	Model interface{}
}

func RegiseterModels() []Models {
	return []Models{
		{Model: User{}},
		{Model: Address{}},
		{Model: Category{}},
		{Model: Order{}},
		{Model: OrderCustomer{}},
		{Model: OrderItem{}},
		{Model: Payment{}},
		{Model: Product{}},
		{Model: ProductImage{}},
		{Model: Section{}},
		{Model: Shipment{}},
	}
}
