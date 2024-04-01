package main

import (
	pb "MyCal/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {

	log.Println("Max server function have been invoked ")
	var max int

	for {

		res, err := stream.Recv()

		if err == io.EOF {
			return err
		}

		if err != nil {
			log.Fatal(err)
		}
		if max < int(res.Number) {
			max = int(res.Number)
			stream.Send(&pb.Maxresponse{
				Result: res.Number,
			})

		}

	}

}
