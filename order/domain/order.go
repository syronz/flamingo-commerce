package domain

type (
	// Order
	Order struct {
		ID           int
		CreationTime string
		UpdateTime   string
		OrderItems   []OrderItem
	}
	// OrderItem
	OrderItem struct {
		Sku          string
		Name         string
		Price        float64
		PriceInclTax float64
	}
)