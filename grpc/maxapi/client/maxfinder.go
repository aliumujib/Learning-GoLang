package main

import (
	"context"
	"github.com/GoLangWebDev/grpc/maxapi/proto"
	"io"
	"log"
	"time"
)

func RequestMax(client proto.MaxAPIServiceClient) {

	stream, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating client stream: %v\n", err)
	}

	requests := []*proto.MaxApiRequest{
		{Number: 2},
		{Number: 1},
		{Number: 5},
		{Number: 6},
		{Number: 3},
	}

	waitChan := make(chan struct{})

	//start a goroutine for sending numbers
	go func() {
		for _, req := range requests {
			stream.Send(req)

			//delay for a second
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	//start another goroutine for receiving the numbers
	go func() {
		for {
			result, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("A serious error occurred while recieving data from server: %v\n", err)
				break
			}

			log.Printf("Received max for stream: %v\n", result.Max)
		}

		close(waitChan)
	}()

	<-waitChan
}
