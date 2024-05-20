package main

import (
	"log"
	"net/http"

	handler "github.con/briheet/kitchen/services/orders/handler/orders"
	"github.con/briheet/kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	ordersService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(ordersService)
	orderHandler.RegisterRouter(router)

	log.Println("starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
