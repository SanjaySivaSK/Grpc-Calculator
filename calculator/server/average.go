package main

import (
	pb "MyCal/calculator/proto"
	"fmt"
	"io"
	"log"
	
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {

	fmt.Println("Average function has been invoked ")

	var res float64
	var count int

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float32(res) / float32(count),
			})
		}

		if err != nil {
			log.Fatal(err)
		}

		res += float64(req.Number)
		count++

	}

}
