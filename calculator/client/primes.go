package main

import (
	pb "MyCal/calculator/proto"
	"context"

	"io"
	"log"

	"google.golang.org/grpc/status"
)

func Primes(c pb.CalculatorServiceClient) {

	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{
		Number: 102,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {

			log.Println("error message in ", e.Message())
			log.Println("error code in ", e.Code())

		} else {
			log.Fatal(err)
		}
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break

		}

		if err != nil {
			log.Fatal(err)

		}
		log.Println("Number are", msg.Result)
	}

}
