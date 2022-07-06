package main

import (
	pb "github.com/GoLangWebDev/grpc/primesapi/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addressString = "0.0.0.0:8080"

type Server struct {
	pb.PrimesAPIServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addressString)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %v\n", addressString)

	server := grpc.NewServer()
	pb.RegisterPrimesAPIServiceServer(server, &Server{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve on: %v\n", err)
	}
}
