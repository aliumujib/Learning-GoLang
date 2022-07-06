package main

import (
	"context"
	"fmt"
	pb "github.com/GoLangWebDev/grpc/primesapi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

//protoc -Iproto --go_opt=module=github.com/GoLangWebDev/grpc/primes_api --go_out=. --go-grpc_opt=module=github.com/GoLangWebDev/grpc/primes_api --go-grpc_out=. proto/*.proto
var addressString = "0.0.0.0:8080"

func main() {
	conn, err := grpc.Dial(addressString, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewPrimesAPIServiceClient(conn)

	stream, err := client.Primes(context.Background(), &pb.PrimesApiRequest{
		Number: 120,
	})

	for {
		msf, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream %v\n", err)
		}

		fmt.Println("Printing numbers: ", msf.Prime)
	}
}
