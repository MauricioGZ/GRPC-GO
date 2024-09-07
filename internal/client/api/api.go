package api

import (
	"log"
	"net/http"

	"github.com/MauricioGZ/GRPC-GO/internal/client/service"
)

type api struct {
	serv service.Service
}

func New(_serv service.Service) *api {
	return &api{
		serv: _serv,
	}
}

func (a *api) Run() error {
	router := http.NewServeMux()
	a.registerRoutes(router)

	log.Println("Starting server on ", "9000")
	return http.ListenAndServe(":9000", router)
}
