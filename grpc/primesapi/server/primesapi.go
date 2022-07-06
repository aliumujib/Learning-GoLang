package main

import (
	"github.com/GoLangWebDev/grpc/primesapi/proto"
)

func (server *Server) Primes(in *proto.PrimesApiRequest, stream proto.PrimesAPIService_PrimesServer) error {
	current := 2
	number := int(in.Number)

	for number > 1 {
		if number%current == 0 {
			stream.Send(&proto.PrimesApiResponse{
				Prime: int32(current),
			})
			number = number / current
		} else {
			current = current + 1
		}
	}

	return nil
}
