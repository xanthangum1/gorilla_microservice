package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xanthangum1/gorilla_micro/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//GetProducts returns products from data store
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// Creates a Product
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

//UpdateProducts changes info with id
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT products", id)

	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}
