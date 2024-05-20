package types

import (
	"context"

	"github.con/briheet/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
