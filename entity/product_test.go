package entity_test

import (
	"testing"

	"product-ctc/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := entity.NewProduct("TEST-01", "Product Test 1", 0, 149900, 100)
	assert.Nil(t, err)
	assert.Equal(t, p.Name, "Product Test 1")
	assert.Equal(t, p.SKU, "TEST-01")
	assert.NotNil(t, p.ID)
}

func TestProductValidate(t *testing.T) {
	type test struct {
		sku       string
		name      string
		catalogID int64
		price     float32
		quantity  int
		message   string
		want      error
	}

	tests := []test{
		{
			sku:       "",
			name:      "Product Test 1",
			catalogID: 0,
			price:     149900,
			quantity:  100,
			want:      nil,
		},
		{
			sku:       "TEST-01",
			name:      "",
			catalogID: 0,
			price:     149900,
			quantity:  100,
			want:      entity.ErrInvalidEntity,
		},
		{
			sku:       "TEST-01",
			name:      "Product Test 1",
			catalogID: 0,
			price:     149900,
			quantity:  -1,
			want:      entity.ErrInvalidEntity,
		},
		{
			sku:       "TEST-01",
			name:      "Product Test 1",
			catalogID: 0,
			price:     1,
			quantity:  100,
			want:      entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewProduct(tc.sku, tc.name, tc.catalogID, tc.price, tc.quantity)
		assert.Equal(t, err, tc.want)
	}

}
