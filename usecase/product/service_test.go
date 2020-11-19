package product

import (
	"testing"
	"time"

	"product-ctc/entity"
	"product-ctc/pkg"

	repository "product-ctc/infrastructure/repository/memcache"

	"github.com/stretchr/testify/assert"
)

func newFixtureProduct() *entity.Product {
	return &entity.Product{
		Name:      "Product Test 1",
		SKU:       "TEST-01",
		CatalogID: 0,
		Price:     149900,
		Quantity:  100,
		Status:    pkg.StatusActive,
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := repository.NewProductMemCache()
	m := NewService(repo)
	u := newFixtureProduct()
	_, err := m.CreateProduct(u.SKU, u.Name, u.CatalogID, u.Price, u.Quantity)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := repository.NewProductMemCache()
	m := NewService(repo)
	u1 := newFixtureProduct()
	u2 := newFixtureProduct()
	u2.Name = "Product Coba 2"

	uID, _ := m.CreateProduct(u1.SKU, u1.Name, u1.CatalogID, u1.Price, u1.Quantity)
	_, _ = m.CreateProduct(u2.SKU, u2.Name, u2.CatalogID, u2.Price, u2.Quantity)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchProducts("test")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Product Test 1", c[0].Name)

		c, err = m.SearchProducts("sandal")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListProducts()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetProduct(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Name, saved.Name)
	})
}

func Test_Update(t *testing.T) {
	repo := repository.NewProductMemCache()
	m := NewService(repo)
	u := newFixtureProduct()
	id, err := m.CreateProduct(u.SKU, u.Name, u.CatalogID, u.Price, u.Quantity)
	assert.Nil(t, err)

	t.Run("update name", func(t *testing.T) {
		saved, _ := m.GetProduct(id)
		saved.Name = "Product Test 2"
		assert.Nil(t, m.UpdateProduct(saved))
		updated, err := m.GetProduct(id)
		assert.Nil(t, err)
		assert.Equal(t, "Product Test 2", updated.Name)
	})

	t.Run("update status", func(t *testing.T) {
		saved, _ := m.GetProduct(id)
		saved.Status = 0
		assert.Nil(t, m.UpdateProduct(saved))
		updated, err := m.GetProduct(id)
		assert.Nil(t, err)
		assert.Equal(t, pkg.StatusInActive, updated.Status)
	})

}

func TestDelete(t *testing.T) {
	repo := repository.NewProductMemCache()
	m := NewService(repo)
	u1 := newFixtureProduct()
	u2 := newFixtureProduct()
	u2ID, _ := m.CreateProduct(u2.SKU, u2.Name, u2.CatalogID, u2.Price, u2.Quantity)

	err := m.DeleteProduct(u1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteProduct(u2ID)
	assert.Nil(t, err)
	_, err = m.GetProduct(u2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
