package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"log"

	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Println("create fucntion have invoked")

	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error in Blog in database ")

	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {

		return nil, status.Error(codes.Internal, "cannot connect to oid ")

	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil

}
