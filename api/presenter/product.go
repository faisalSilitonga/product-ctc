package presenter

import (
	"product-ctc/entity"
)

//Product data
type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	CatalogID int64     `json:"catalogid,omitempty"`
	Price     float32   `json:"price"`
	Quantity  int       `json:"quantity,omitempty"`
	Status    string    `json:"status,omitempty"`
}
