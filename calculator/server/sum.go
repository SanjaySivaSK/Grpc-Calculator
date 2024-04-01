package main

import (
	pb "MyCal/calculator/proto"
	"context"
)

func (s *Server) Sum( ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error){
	return &pb.SumResponse{
      Result: in.Number1+in.Number2,
	},nil
}
