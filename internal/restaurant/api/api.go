package api

import (
	"log"
	"net/http"

	sr "github.com/MauricioGZ/GRPC-GO/internal/restaurant/service"
)

type api struct {
	serv sr.Service
}

func New(_serv sr.Service) *api {
	return &api{
		serv: _serv,
	}
}

func (a *api) Run() error {
	router := http.NewServeMux()
	a.registerRoutes(router)

	log.Println("Starting server on ", "9000")
	return http.ListenAndServe(":9090", router)
}
