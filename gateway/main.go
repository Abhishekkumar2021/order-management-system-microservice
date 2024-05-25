package main

import (
	"log"
	"net/http"

	"github.com/Abhishekkumar2021/commons"
	pb "github.com/Abhishekkumar2021/commons/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddress      = commons.EnvStrings("HTTP_ADDRESS", ":3000")
	orderServiceAddr = commons.EnvStrings("ORDER_SERVICE_ADDR", "localhost:50051")
)

func main() {
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(client)
	handler.registerRoutes(mux)
	log.Printf("Server listening on %s", httpAddress)
	if err := http.ListenAndServe(httpAddress, mux); err != nil {
		panic(err)
	}
}
