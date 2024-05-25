package main

import (
	"context"

	pb "github.com/Abhishekkumar2021/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	// grpc service
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server) *grpcHandler {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
	return handler
}

func (h *grpcHandler) CreateOrder(context context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// read data from req
	customerID := req.CustomerID

	// return response
	response := &pb.CreateOrderResponse{
		OrderID:    "123",
		CustomerID: customerID,
	}
	return response, nil
}
