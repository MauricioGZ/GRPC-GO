package api

import (
	"context"
	"net/http"
	"time"

	"github.com/MauricioGZ/GRPC-GO/internal/utils"
)

func (a *api) getAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	products, err := a.serv.GetAllProducts(ctx)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.JSONAnyResponse(w, http.StatusOK, products)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}
