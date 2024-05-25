package main

import (
	"context"
	"log"
	"net"

	"github.com/Abhishekkumar2021/commons"
	"google.golang.org/grpc"
)

var grpcAddress = commons.EnvStrings("GRPC_ADDRESS", "localhost:50051")

func main() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	store := NewStore()
	service := NewService(store)
	service.CreateOrder(context.Background())
	log.Printf("grpc server started at %s", grpcAddress)
	NewGrpcHandler(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}
