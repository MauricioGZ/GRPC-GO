package model

type Product struct {
	ProductID uint32  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
}
