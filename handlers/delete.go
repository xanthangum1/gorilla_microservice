package handlers

import (
	"net/http"

	"github.com/xanthangum1/gorilla_micro/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns nothing
// responses:
//	201: noContent

// DeleteProduct deletes a product from data store
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting recordid", id)

	err := data.ErrProductNotFound(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
