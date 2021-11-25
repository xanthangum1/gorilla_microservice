package handlers

import (
	"net/http"

	"github.com/xanthangum1/gorilla_micro/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products from the database
// responses:
//	200: productsReponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

//swagger:route GET /products/{id} products listSingle
// Reutnr a list of proucts from the database
// responses:
// 	200: productResponse
//  404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetPRoductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(Http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod. rw)
	if err != nil {
		// if error on making response
		p.l.Println("[ERROR] serializing product", err)
	}

}