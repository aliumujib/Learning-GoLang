package main

import (
	"context"
	pb "github.com/GoLangWebDev/grpc/calculator/proto"
)

func (calculator *Server) Sum(_ context.Context, request *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Sum: request.FirstNumber + request.SecondNumber,
	}, nil
}
