package views

import (
	"encoding/json"
	"github.com/artemmarkaryan/gotints/internal/domain"
	"log"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := domain.GetProducts(r.Context())
	if err != nil {
		log.Panic("cant get products: ", err)
	}

	response, err := json.Marshal(products)
	if err != nil {
		log.Panic("cant serialize products: ", err)
	}

	_, err = w.Write(response)
	if err != nil {
		log.Panic("cant write response: ", err)
	}
}
