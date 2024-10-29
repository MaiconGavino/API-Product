package main

import (
	"apiproducts/src/pb/products"
	"apiproducts/src/repository"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	products.ProductServiceServer
	productRepo *repository.ProductRepository
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	newProduct, err := s.productRepo.Create(*product)
	if err != nil {
		return nil, err
	}
	return &newProduct, nil
}

func (s *server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	productList, err := s.productRepo.FindAll()
	return &productList, err
}

func main() {

	srv := server{}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("error on create listerner. erro:", err)
	}

	s := grpc.NewServer()
	products.RegisterProductServiceServer(s, &srv)

	if err := s.Serve(listener); err != nil {
		log.Fatalln("error on serve. erro:", err)
	}
}
