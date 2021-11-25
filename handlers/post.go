package handlers

import (
	"net/http"

	"github.com/xanthangum1/gorilla_micro/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
// 		200: productReponse
// 422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// featch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#c\n", prod)
	data.Addproduct(prod)
}