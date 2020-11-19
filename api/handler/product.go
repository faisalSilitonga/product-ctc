package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"product-ctc/api/presenter"
	"product-ctc/entity"
	"product-ctc/pkg"
	"product-ctc/usecase/product"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listProducts(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading Products"
		var data []*entity.Product
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.ListProducts()
		default:
			data, err = service.SearchProducts(name)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Product
		for _, d := range data {
			toJ = append(toJ, &presenter.Product{
				ID:    d.ID,
				Name:  d.Name,
				SKU:   d.SKU,
				Price: d.Price,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding Product"
		var input struct {
			Name      string  `json:"name"`
			SKU       string  `json:"sku"`
			CatalogID int64   `json:"catalogid"`
			Price     float32 `json:"price"`
			Quantity  int     `json:"quantiry"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateProduct(input.SKU, input.Name, input.CatalogID, input.Price, input.Quantity)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Product{
			ID:        id,
			Name:      input.Name,
			SKU:       input.SKU,
			CatalogID: input.CatalogID,
			Price:     input.Price,
			Quantity:  input.Quantity,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding Product"
		var input struct {
			ID        string  `json:"id"`
			Name      string  `json:"name"`
			SKU       string  `json:"sku"`
			CatalogID int64   `json:"catalogid"`
			Price     float32 `json:"price"`
			Quantity  int     `json:"quantiry"`
			Status    int     `json:"status"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := entity.StringToID(input.ID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		product := entity.Product{
			ID:        id,
			Name:      input.Name,
			SKU:       input.SKU,
			CatalogID: input.CatalogID,
			Price:     input.Price,
			Quantity:  input.Quantity,
			Status:    input.Status,
		}

		err = service.UpdateProduct(&product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Product{
			ID:        product.ID,
			Name:      input.Name,
			SKU:       input.SKU,
			CatalogID: input.CatalogID,
			Price:     input.Price,
			Quantity:  input.Quantity,
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading Product"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetProduct(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		status := pkg.StrStatusActive
		if data.Status == pkg.StatusInActive {
			status = pkg.StrStatusInActive
		}

		toJ := &presenter.Product{
			ID:        data.ID,
			Name:      data.Name,
			SKU:       data.SKU,
			CatalogID: data.CatalogID,
			Price:     data.Price,
			Quantity:  data.Quantity,
			Status:    status,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deleteProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing Product"
		vars := mux.Vars(r)

		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.DeleteProduct(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.Write([]byte("OK"))
	})
}

//MakeProductHandlers make url handlers
func MakeProductHandlers(r *mux.Router, n negroni.Negroni, service product.UseCase) {
	r.Handle("/v1/product", n.With(
		negroni.Wrap(listProducts(service)),
	)).Methods("GET", "OPTIONS").Name("listProducts")

	r.Handle("/v1/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS").Name("getProduct")

	r.Handle("/v1/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS").Name("createProduct")

	r.Handle("/v1/product/update", n.With(
		negroni.Wrap(updateProduct(service)),
	)).Methods("POST", "OPTIONS").Name("updateProduct")

	r.Handle("/v1/product/{id}", n.With(
		negroni.Wrap(deleteProduct(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteProduct")
}
