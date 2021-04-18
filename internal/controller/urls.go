package controller

import (
	"github.com/artemmarkaryan/gotints/internal/views"
	"net/http"
)

func ConfigureRoutes() {
	http.HandleFunc("/product/all", views.GetProducts)
}
