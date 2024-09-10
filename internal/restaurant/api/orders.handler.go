package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (a *api) GetPendingOrders(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := a.serv.GetPendingOrders(ctx)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
