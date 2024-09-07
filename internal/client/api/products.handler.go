package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (a *api) getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	products, err := a.serv.GetAllProducts(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Fatal(err)
	}
}
