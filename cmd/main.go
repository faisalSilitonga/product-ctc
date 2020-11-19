package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	repository "product-ctc/infrastructure/repository/memcache"

	"product-ctc/api/handler"
	"product-ctc/api/middleware"
	"product-ctc/config"
	"product-ctc/usecase/product"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {

	productRepo := repository.NewProductMemCache()
	productService := product.NewService(productRepo)

	r := mux.NewRouter()

	//register middleware
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	//register product handler
	handler.MakeProductHandlers(r, *n, productService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PONG"))
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
