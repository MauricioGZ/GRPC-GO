package api

import (
	"context"
	"net/http"
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
