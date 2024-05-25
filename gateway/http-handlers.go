package main

import (
	"net/http"

	"github.com/Abhishekkumar2021/commons"
	pb "github.com/Abhishekkumar2021/commons/api"
)

type handler struct {
	// gateway service
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client: client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	// Add routes here
	mux.HandleFunc("POST /api/customers/{id}/orders", h.createOrder)
}

func (h *handler) createOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("id")
	var items []*pb.ItemsWithQuantity
	if err := commons.ReadJSON(r, &items); err != nil {
		commons.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	// Call the gRPC service
	response, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})

	if err != nil {
		commons.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	commons.WriteJSON(w, http.StatusOK, response)
}
