package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"product-ctc/entity"
	"product-ctc/usecase/product/mock"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_listProducts(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("listProducts").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/product", path)
	b := &entity.Product{
		ID: entity.NewID(),
	}
	service.EXPECT().
		ListProducts().
		Return([]*entity.Product{b}, nil)
	ts := httptest.NewServer(listProducts(service))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listProducts_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listProducts(service))
	defer ts.Close()
	service.EXPECT().
		SearchProducts("product of products").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?name=product+of+products")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listProducts_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	b := &entity.Product{
		ID: entity.NewID(),
	}
	service.EXPECT().
		SearchProducts("ozzy").
		Return([]*entity.Product{b}, nil)
	ts := httptest.NewServer(listProducts(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?name=ozzy")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("createProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/product", path)

	service.EXPECT().
		CreateProduct(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil)
	h := createProduct(service)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
		"name": "Product Test 1",
		"sku": "TEST-01",
		"catalogid": 294,
		"price":100,
		"quantity":1
		}`)
	resp, _ := http.Post(ts.URL+"/v1/product", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var p *entity.Product
	json.NewDecoder(resp.Body).Decode(&p)
	assert.Equal(t, "Product Test 1", p.Name)
}

func Test_updateProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("updateProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/product/update", path)

	service.EXPECT().UpdateProduct(gomock.Any()).Return(nil)

	handler := updateProduct(service)
	r.Handle("/v1/product/update", handler).Methods("POST", "OPTIONS")
	ts := httptest.NewServer(r)
	defer ts.Close()

	payload := fmt.Sprintf(`{
		"id": "151a6b1a-8dfa-45a2-b29c-969d4755224f",
		"name": "Product Test 1",
		"sku": "TEST-01",
		"catalogid": 294,
		"price":100,
		"quantity":1,
		"status":1
		}`)

	resp, err := http.Post(ts.URL+"/v1/product/update", "application/json", strings.NewReader(payload))
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var p *entity.Product
	json.NewDecoder(resp.Body).Decode(&p)
	assert.Equal(t, "Product Test 1", p.Name)
}

func Test_getProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("getProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/product/{id}", path)
	b := &entity.Product{
		ID: entity.NewID(),
	}
	service.EXPECT().
		GetProduct(b.ID).
		Return(b, nil)
	handler := getProduct(service)
	r.Handle("/v1/product/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/product/" + b.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *entity.Product
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, b.ID, d.ID)
}

func Test_deleteProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("deleteProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/product/{id}", path)
	b := &entity.Product{
		ID: entity.NewID(),
	}
	service.EXPECT().DeleteProduct(b.ID).Return(nil)
	handler := deleteProduct(service)
	req, _ := http.NewRequest("DELETE", "/v1/product/"+b.ID.String(), nil)
	r.Handle("/v1/productmark/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
