package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MauricioGZ/GRPC-GO/internal/client/api/dto"
	"github.com/MauricioGZ/GRPC-GO/internal/utils"
)

func (a *api) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderItems []dto.OrderItem
	err := utils.ParseJSON(r, &orderItems)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	order, err := a.serv.CreateOrder(ctx, orderItems)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.JSONAnyResponse(w, http.StatusOK, order)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (a *api) CancelOrder(w http.ResponseWriter, r *http.Request) {
	orderIDString := r.PathValue("id")

	orderID, err := strconv.ParseUint(orderIDString, 10, 32)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = a.serv.CancelOrder(ctx, uint32(orderID))
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSONMessageResponse(w, http.StatusOK, fmt.Sprintf("order with orderID: %d canceled", orderID))
}
