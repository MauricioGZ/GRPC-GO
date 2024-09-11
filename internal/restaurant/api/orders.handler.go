package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MauricioGZ/GRPC-GO/internal/utils"
)

func (a *api) GetPendingOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	pendingOrders, err := a.serv.GetPendingOrders(ctx)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.JSONAnyResponse(w, http.StatusOK, pendingOrders)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (a *api) SetOrderToReady(w http.ResponseWriter, r *http.Request) {
	orderIDString := r.PathValue("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	orderID, err := strconv.ParseUint(orderIDString, 10, 32)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	err = a.serv.SetOrderToReady(ctx, uint32(orderID))

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSONMessageResponse(w, http.StatusOK, fmt.Sprintf("order_id: %d set to ready", orderID))
}
