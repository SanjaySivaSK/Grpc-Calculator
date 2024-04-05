package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("listofblog function have invoked ")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		status.Error(codes.Internal, "internal error occur in server")
	}
    defer cur.Close(context.Background())


    for cur.Next(context.Background()){
		data:=&BlogItem{}
		err:=cur.Decode(data)// here we setting the field with mongodb datas with data of blogitem(mongo)
		// the decode do the work of data from mongo to data blogitems 
		if err != nil {
			log.Fatal(err)
		}

		stream.Send(documentToBlog(data))// heret hte 
	}


	
	return nil
}
