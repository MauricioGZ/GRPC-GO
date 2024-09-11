package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func (a *api) SetOrderToReady(w http.ResponseWriter, r *http.Request) {
	orderIDString := r.PathValue("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	orderID, err := strconv.ParseUint(orderIDString, 10, 32)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = a.serv.SetOrderToReady(ctx, uint32(orderID))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
