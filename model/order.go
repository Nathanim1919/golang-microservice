package model

import (
	"github.com/google/uuid"
)

type Order struct {
	OrderID uint64          `jsoon:"order_id"`
	CustomerID uuid.UUID	`jsoon:"customer_id"`
	LineItems  []LineItem	`jsoon:"line_items"`
	CreatedAt  *time.Time	`jsoon:"created_at"`
	ShippedAt  *time.Time	`jsoon:"shipped_at"`
	CompletedAt  *time.Time	`jsoon:"completed_at"`
}

type LineItem struct {
	ItemID uuid.UUID `jsoon:"item_id"`
	Quantity uint	 `jsoon:"quantity"`
	Price    uint	 `jsoon:"price"`
}
