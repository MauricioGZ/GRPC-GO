package api

import (
	"net/http"
)

func (a *api) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /order", a.GetPendingOrders)
	router.HandleFunc("PATCH /order/{id}", a.SetOrderToReady)
}
