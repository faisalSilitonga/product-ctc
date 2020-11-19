package repository

import (
	"database/sql"
	"time"

	"product-ctc/entity"
)

//ProductMySQL mysql repo
type ProductMySQL struct {
	db *sql.DB
}

//NewProductMySQL create new repository
func NewProductMySQL(db *sql.DB) *ProductMySQL {
	return &ProductMySQL{
		db: db,
	}
}

//Create a Product
func (r *ProductMySQL) Create(e *entity.Product) (entity.ID, error) {

	// LOGIC INSERT DATA TO MYSQL

	return e.ID, nil
}

//Get a Product
func (r *ProductMySQL) Get(id entity.ID) (*entity.Product, error) {

	var p entity.Product

	// LOGIC GET DATA

	return &p, nil
}

//Update a Product
func (r *ProductMySQL) Update(e *entity.Product) error {
	e.UpdatedAt = time.Now()

	// LOGIC UPDATE DATA

	return nil
}

//Search Products
func (r *ProductMySQL) Search(query string) ([]*entity.Product, error) {
	var products []*entity.Product

	// LOGIC GET LIST BY QUERY DATA

	return products, nil
}

//List Products
func (r *ProductMySQL) List() ([]*entity.Product, error) {
	var products []*entity.Product

	// LOGIC GET LIST DATA

	return products, nil
}

//Delete a Product
func (r *ProductMySQL) Delete(id entity.ID) error {
	// LOGIC DELETE DATA
	return nil
}
