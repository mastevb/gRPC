package main

import (
	"context"
	"errors"
	"log"

	"github.com/gofrs/uuid"
	pb "productinfo/service/ecommerce"
)

// server is used to implement ecommerce/product_info
type server struct {
	productMap map[string]*pb.Product
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {

}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {

}
