package handlers

import (
	"net/http"

	"github.com/xanthangum1/gorilla_micro/data"
)

// swagger:route GET /products products GetProducts
// Returns a list of products
// responses:
//	200: productsReponse

// GetProducts return the poducts from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {

	//fetch products from datastore
	lp := data.GetProducts()

	//serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
