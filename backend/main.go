package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/radityaqb/tgtc/backend/handlers"
	"github.com/radityaqb/tgtc/backend/server"
)

type Product struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	ShopName     string  `json:"shop_name"`
	ProductPrice float64 `json:"product_price"`
	ImageURL     string  `json:"image_url"`
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error_message"`
}

var products = []Product{
	Product{
		ID:           1,
		Name:         "Nama1",
		ShopName:     "Tokopedia",
		ProductPrice: 666.0,
		ImageURL:     "www.google.com",
	},
}
var APIResponses = []APIResponse{}

func main() {

	// Init database connection
	// database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	// routes http
	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)

	// construct your own API endpoints
	// endpoint : /add-product
	router.HandleFunc("/add-product", addProduct).Methods(http.MethodPost)
	// endpoint : /get-product?id=
	router.HandleFunc("/get-product/{id}", getProduct).Methods(http.MethodGet)
	// endpoint : /update-product
	// endpoint : /delete-product

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, product := range products {
		dump, _ := strconv.ParseInt(params["id"], 10, 64)
		if product.ID == dump {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	product := Product{}
	_ = json.NewDecoder(r.Body).Decode(&product)
	// product.ID = strconv.Itoa(len(products) + 1)
	product.ID = int64(len(products) + 1)
	products = append(products, product)

	json.NewEncoder(w).Encode(product)
}
