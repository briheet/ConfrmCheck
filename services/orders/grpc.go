package main

import (
	"log"
	"net"

	handler "github.con/briheet/kitchen/services/orders/handler/orders"
	"github.con/briheet/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("failed to listen: %v", s.addr)
	}

	grpcServer := grpc.NewServer()

	// register grpc services

	orderService := service.NewOrderService()
	handler.NewGrpcService(grpcServer, orderService)

	log.Println("starting a new gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
