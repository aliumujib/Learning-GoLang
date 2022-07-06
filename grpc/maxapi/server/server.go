package main

import (
	pb "github.com/GoLangWebDev/grpc/maxapi/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addressString = "0.0.0.0:8080"

type Server struct {
	pb.MaxAPIServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addressString)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %v\n", addressString)

	server := grpc.NewServer()
	pb.RegisterMaxAPIServiceServer(server, &Server{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve on: %v\n", err)
	}
}
