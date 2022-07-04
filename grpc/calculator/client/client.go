package main

import (
	"context"
	"fmt"
	pb "github.com/GoLangWebDev/grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//protoc -Iproto --go_opt=module=github.com/GoLangWebDev/grpc/calculator --go_out=. --go-grpc_opt=module=github.com/GoLangWebDev/grpc/calculator --go-grpc_out=. proto/*.proto
var addressString = "0.0.0.0:8080"

func main() {
	conn, err := grpc.Dial(addressString, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	sum, err := client.Sum(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  2,
		SecondNumber: 3,
	})

	fmt.Println("Sum of the operation is: ", sum)
}
