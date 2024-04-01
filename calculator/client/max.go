package main

import (
	pb "MyCal/calculator/proto"
	"context"
	"io"
	"log"
	"time"
)

func Domax(c pb.CalculatorServiceClient) {

	log.Println("Max function have been invoked ")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Println(" error in max fucntion of client side ")
	}

	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
		


	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			// log.Print("SENDING:" ,req)
			stream.Send((*pb.MaxRequest)(req))
			time.Sleep(1*time.Second)

		}
		stream.CloseSend()
	}()

	go func() {

		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				// log.Println("err occur in the receive of response ")
				break

			}

			log.Println("max no are", res.Result)

		}
		close(waitc)

	}()

	<-waitc

}
