package entity

import "time"

type Order struct {
	ID         uint32
	CustomerID uint32
	OrderDate  time.Time
	Status     string
	TotalPrice float32
}
