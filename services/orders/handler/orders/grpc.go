package handler

import (
	"context"

	"github.con/briheet/kitchen/services/common/genproto/orders"
	"github.con/briheet/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	// service injection
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcService(orpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{}
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}