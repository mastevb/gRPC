package main

import (
	"context"
	"log"
	"time"

	pb "productinfo/client/ecommerce"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// create a unsecured connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)
	// add some products
	name := "Apple iPhone 12"
	description := "Just another iPhone"
	price := float32(1000.0)
	// Context contains metadata
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// invoke rpc call for add
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Cloud not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)
	// invoke rpc call for get
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %v", product.String())
}