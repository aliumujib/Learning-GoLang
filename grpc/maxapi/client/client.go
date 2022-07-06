package main

import (
	pb "github.com/GoLangWebDev/grpc/maxapi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//protoc -Iproto --go_opt=module=github.com/GoLangWebDev/grpc/max_api --go_out=. --go-grpc_opt=module=github.com/GoLangWebDev/grpc/max_api --go-grpc_out=. proto/*.proto
var addressString = "0.0.0.0:8080"

func main() {
	conn, err := grpc.Dial(addressString, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewMaxAPIServiceClient(conn)

	RequestMax(client)
}
