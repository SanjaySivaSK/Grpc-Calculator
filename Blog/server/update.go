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

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {

	log.Println("update blog have been implemented")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "cannot parse the Id")
	}

	data:=&BlogItem{
		AuthorID: in.AuthorId,
		Title: in.Title,
		Content: in.Content,
	}

	res,err:=collection.UpdateOne(ctx,bson.M{"_id":oid},bson.M{"$set":data})

	if err != nil {
		log.Fatal(err)
	}

	if res.MatchedCount==0{
		return nil,err
	}

	return &emptypb.Empty{},nil

}