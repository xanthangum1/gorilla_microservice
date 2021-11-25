package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xanthangum1/gorilla_micro/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns nothing
// responses:
//	201: noContent

// DeleteProduct deletes a product from data store
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	// this will always convert because of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("handle DELETE Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
