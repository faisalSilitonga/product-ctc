package product

import (
	"product-ctc/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Product, error)
	Search(query string) ([]*entity.Product, error)
	List() ([]*entity.Product, error)
}

//Writer Product writer
type Writer interface {
	Create(e *entity.Product) (entity.ID, error)
	Update(e *entity.Product) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetProduct(id entity.ID) (*entity.Product, error)
	SearchProducts(query string) ([]*entity.Product, error)
	ListProducts() ([]*entity.Product, error)
	CreateProduct(sku string, name string, catalogID int64, price float32, quantity int) (entity.ID, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id entity.ID) error
}
