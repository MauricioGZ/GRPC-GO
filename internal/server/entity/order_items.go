package entity

type OrderItem struct {
	ID          uint32
	OrderID     uint32
	ProductID   uint32
	Quantity    uint32
	ProductName string
}
