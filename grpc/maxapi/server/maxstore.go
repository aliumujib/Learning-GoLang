package main

import (
	"github.com/GoLangWebDev/grpc/maxapi/proto"
	"github.com/adam-lavrik/go-imath/ix"
	"io"
	"log"
)

var nums []int

func (server *Server) Max(stream proto.MaxAPIService_MaxServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		nums = append(nums, int(req.Number))
		max := ix.MaxSlice(nums)

		err = stream.Send(&proto.MaxApiResponse{
			Max: int32(max),
		})

		if err != nil {
			log.Fatalf("error while responding to the client: %v\n", err)
		}
	}

	return nil
}
