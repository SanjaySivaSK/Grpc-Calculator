package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("the delete function have invoked",in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "we cannot parse the id ")
	}
	done, err := collection.DeleteOne(ctx,bson.M{"_id":oid})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "error occur in deletion  ")
	}

	if done.DeletedCount == 0 {
		return nil, status.Error(codes.Internal, "delection doesn't done")
	}

	return nil, nil

}
