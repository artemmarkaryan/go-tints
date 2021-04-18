package views

import (
	"github.com/artemmarkaryan/gotints/internal/domain"
	"log"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := domain.GetProducts(r.Context())
	if err != nil {
		log.Panic("cant get products: ", err)
	}

	_, err = w.Write([]byte(products))
	if err != nil {
		log.Panic("cant write response: ", err)
	}
}
