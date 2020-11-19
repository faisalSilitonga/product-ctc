package entity

import (
	"fmt"
	"product-ctc/pkg"
	"time"
)

//Product data
type Product struct {
	ID        ID
	SKU       string
	Name      string
	CatalogID int64
	Price     float32
	Quantity  int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewProduct create a new product
func NewProduct(sku string, name string, catalogID int64, price float32, quantity int) (*Product, error) {
	p := &Product{
		ID:        NewID(),
		SKU:       sku,
		Name:      name,
		CatalogID: catalogID,
		Price:     price,
		Quantity:  quantity,
		Status:    pkg.StatusActive,
		CreatedAt: time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

//Validate validate product
func (p *Product) Validate() error {
	if p.Name == "" || p.Quantity < 0 || p.Price <= 100 {
		return ErrInvalidEntity
	}
	return nil
}

//Show print detail product
func (p *Product) Show() {
	fmt.Println("Name = ", p.Name)
}
