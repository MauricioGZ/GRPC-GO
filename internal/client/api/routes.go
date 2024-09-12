package api

import (
	"net/http"
)

func (a *api) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /products", a.getAllProducts)
	router.HandleFunc("POST /order", a.CreateOrder)
	router.HandleFunc("DELETE /order/{id}", a.CancelOrder)
}
