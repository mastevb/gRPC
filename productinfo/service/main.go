package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "productinfo/service/ecommerce"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// new gRPC server instance is created by calling gRPC Go APIs
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
