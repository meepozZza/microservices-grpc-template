package main

import (
	handler "grpc-microservice/services/orders/handler/orders"
	"grpc-microservice/services/orders/service"
	"log"
	"net/http"
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

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersService(orderService)
	orderHandler.RegisterRouter(router)


	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
