package dto

type OrderItem struct {
	ProductID uint32 `json:"product_id"`
	Quantity  uint32 `json:"quantity"`
}
