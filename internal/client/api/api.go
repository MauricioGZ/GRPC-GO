package api

import (
	"log"
	"net/http"

	cs "github.com/MauricioGZ/GRPC-GO/internal/client/service"
)

type api struct {
	serv cs.Service
}

func New(_serv cs.Service) *api {
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
