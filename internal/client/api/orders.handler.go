package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (a *api) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	orderID, err := a.serv.CreateOrder(ctx)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(struct {
		OrderID uint32 `json:"order_id"`
	}{OrderID: orderID}); err != nil {
		log.Fatal(err)
	}
}
