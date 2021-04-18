package views

import (
	"github.com/artemmarkaryan/gotints/internal/domain"
	"log"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	payload, err := domain.GetProducts(r.Context())
	if err != nil {
		payload = err.Error()
	}

	_, err = w.Write([]byte(payload))

	if err != nil {
		log.Fatal("cant write response: ", err)
	}
}
