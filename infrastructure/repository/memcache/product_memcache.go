package repository

import (
	"strings"
	"sync"

	"product-ctc/entity"
	"product-ctc/pkg"
)

//ProductMemCache in memory repo
type ProductMemCache struct {
	sync.RWMutex
	m map[entity.ID]*entity.Product
}

//newProductMemCache create new repository
func NewProductMemCache() *ProductMemCache {
	var m = map[entity.ID]*entity.Product{}
	return &ProductMemCache{
		m: m,
	}
}

//Create a Product
func (r *ProductMemCache) Create(e *entity.Product) (entity.ID, error) {
	r.RLock()
	r.m[e.ID] = e
	r.RUnlock()
	return e.ID, nil
}

//Get a Product
func (r *ProductMemCache) Get(id entity.ID) (*entity.Product, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a Product
func (r *ProductMemCache) Update(e *entity.Product) error {
	oldData, err := r.Get(e.ID)
	if err != nil {
		return err
	}

	if e.Status != pkg.StatusActive {
		e.Status = pkg.StatusInActive
	}

	e.CreatedAt = oldData.CreatedAt

	r.RLock()
	r.m[e.ID] = e
	r.RUnlock()
	return nil
}

//Search Products
func (r *ProductMemCache) Search(query string) ([]*entity.Product, error) {
	var d []*entity.Product
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List Products
func (r *ProductMemCache) List() ([]*entity.Product, error) {
	var d []*entity.Product
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a Product
func (r *ProductMemCache) Delete(id entity.ID) error {
	_, err := r.Get(id)
	if err != nil {
		return err
	}

	r.RLock()
	delete(r.m, id)
	r.RUnlock()
	return nil
}
