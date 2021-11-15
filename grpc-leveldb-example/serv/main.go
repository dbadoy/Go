package main

import (
	"log"
	"net"

	"context"

	"github.com/syndtr/goleveldb/leveldb"
	pb "protos/KVstore"
	"google.golang.org/grpc"
)

const PORT = "3333"

type server struct {
	pb.KVstoreServer
	Db *leveldb.DB
}

func (s *server) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Get : %v", request.GetKey())
	
	data, err := levelGet(s.Db, request.GetKey())
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{Value: data}, nil
}

func (s *server) Set(ctx context.Context, request *pb.SetRequest) (*pb.SetResponse, error) {
	log.Printf("Set : %v, %v", request.GetKey(), request.GetValue())

	err := levelSet(s.Db, request.GetKey(), request.GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.SetResponse{Key:request.GetKey()}, nil
}

func levelSet(db *leveldb.DB, key string, value string) (err error) {
	err = db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return 
	}
	return nil
}

func levelGet(db *leveldb.DB, key string) (v string, err error) {
	var data []byte
	data, err = db.Get([]byte(key), nil) 
	if err != nil {
		return "", err
	}
	return decode(data), nil
}

func decode(b []byte) string {
	return string(b[:len(b)])
}

func main() {
	db, err := leveldb.OpenFile("./level", nil) 
	if err != nil {
		log.Fatalf("failed open leveldb : %s", err)
	}

	defer db.Close()

	listen, err := net.Listen("tcp", ":" + PORT)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterKVstoreServer(grpcServer, &server{Db: db})
	log.Printf("start gRPC server on %s port", PORT)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve : %s", err)
	}
}
