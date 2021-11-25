package handlers

import (
	"net/http"

	"github.com/xanthangum1/gorilla_micro/data"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {

	// fetch the product form the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == dta.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Productnot found in datavase"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
