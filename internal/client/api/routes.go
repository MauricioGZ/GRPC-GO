package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	pb "github.com/MauricioGZ/GRPC-GO/internal/gen"
)

func (a *api) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		menu, err := a.client.GetMenu(ctx, &pb.GetMenuRequest{})

		if err != nil {
			log.Fatal(err)
		}

		done := make(chan bool)

		go func() {
			for {
				res, err := menu.Recv()
				if err != nil {
					if err == io.EOF {
						done <- true
						return
					}
					log.Fatal(err)
				}
				fmt.Println(res.Product)
			}
		}()
		<-done
	})
}
