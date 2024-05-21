package service

import (
	"context"

	"github.con/briheet/kitchen/services/common/genproto/orders"
)

var ordersDB = make([]*orders.Order, 0)

type OrderService struct {
	// orders
	// dependency ingection
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDB = append(ordersDB, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDB
}
