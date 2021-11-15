package main

import (
	"log"
	"time"

	"context"

	pb "protos/KVstore"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3333", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewKVstoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	key := "peer1"
	value := "192.168.0.1"
	r1, err := c.Set(ctx, &pb.SetRequest{Key: key, Value: value})
	if err != nil {
		log.Fatalf("could not request : %v", err)
	}
	log.Printf("result : %v", r1)

	r2, err := c.Get(ctx, &pb.GetRequest{Key: key})
	if err != nil {
		log.Fatalf("could not request : %v", err)
	}
	log.Printf("result : %v", r2)
	
}