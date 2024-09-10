package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (a *api) GetPendingOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	pendingOrders, err := a.serv.GetPendingOrders(ctx)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pendingOrders); err != nil {
		log.Fatal(err)
	}
}
