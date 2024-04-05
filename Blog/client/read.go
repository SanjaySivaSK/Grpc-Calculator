package main

import (
	pb "MyCal/Blog/proto"
	"context"
	"log"
)

func ReadBlog(c pb.BlogServiceClient, id string) *pb.Blog {

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
	return &pb.Blog{
		Id:       res.Id,
		AuthorId: res.AuthorId,
		Title:    res.Title,
		Content:  res.Content,
	}

}
