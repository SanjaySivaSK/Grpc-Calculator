package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {

	log.Println("Read the blog")
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s, error of server while oid", err)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	result := collection.FindOne(ctx, filter)
	err = result.Decode(data)

	if err != nil {
		log.Fatal(err, "error occur in decoding")

	}

	return documentToBlog(data), nil

}
