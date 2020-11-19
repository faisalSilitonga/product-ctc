package product

import (
	"strings"
	"time"

	"product-ctc/entity"
)

//Service product usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//GetProduct get a product
func (s *Service) GetProduct(id entity.ID) (*entity.Product, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, entity.ErrNotFound
	}

	return b, nil
}

//SearchProducts search product
func (s *Service) SearchProducts(query string) ([]*entity.Product, error) {
	products, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}
	return products, nil
}

//ListProducts list products
func (s *Service) ListProducts() ([]*entity.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}
	return products, nil
}

//CreateProduct create a product
func (s *Service) CreateProduct(sku string, name string, catalogID int64, price float32, quantity int) (entity.ID, error) {
	p, err := entity.NewProduct(sku, name, catalogID, price, quantity)
	if err != nil {
		return entity.NewID(), err
	}
	return s.repo.Create(p)
}

//UpdateProduct Update a product
func (s *Service) UpdateProduct(e *entity.Product) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

//DeleteProduct Delete a product
func (s *Service) DeleteProduct(id entity.ID) error {
	_, err := s.GetProduct(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
