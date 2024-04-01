package main

import (
		pb "MyCal/calculator/proto"
	"context"
)

func (s *Server) Div(ctx context.Context, in *pb.DivRequest) (*pb.DivResponse, error) {
	return &pb.DivResponse{
		Quotient: in.Number1 / in.Number2,
		Reminder: in.Number1 % in.Number2,  
	}, nil
}
