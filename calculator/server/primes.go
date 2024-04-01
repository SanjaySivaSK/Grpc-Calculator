package main

import (
	pb "MyCal/calculator/proto"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {

	log.Println("primes function have been invoke ")
	res := in.Number
	k := 2

	if in.Number < 0 {

		return status.Error(codes.InvalidArgument, fmt.Sprintln("The Negative value cannot be processed "))
	} else {
		for res > 1 {
			if res%int64(k) == 0 {
				stream.Send(&pb.PrimesResponse{
					Result: int64(k),
				})
				res = res / int64(k)

			} else {
				k = k + 1

			}

		}
	}
	return nil

}
