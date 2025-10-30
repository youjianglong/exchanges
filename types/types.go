package types

import (
	"strings"
)

// OrderSide 订单方向
type OrderSide string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)

func (o OrderSide) Equal(other string) bool {
	return o == OrderSide(strings.ToLower(other))
}

// OrderStatus 订单状态
type OrderStatus string

const (
	Created         OrderStatus = "created"          // 已创建
	Filled          OrderStatus = "filled"           // 已成交
	Cancelled       OrderStatus = "cancelled"        // 已取消
	Failed          OrderStatus = "failed"           // 失败
	PartiallyFilled OrderStatus = "partially_filled" // 部分成交
	Expired         OrderStatus = "expired"          // 已过期
	Unknown         OrderStatus = "unknown"          // 未知
)

type PosSide string

const (
	PosSideLong  PosSide = "LONG"
	PosSideShort PosSide = "SHORT"
)

type OrderType string

func (o OrderType) Equal(other string) bool {
	return o == OrderType(strings.ToLower(other))
}

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

const (
	IncomeTypeRealizedPnl = "REALIZED_PNL"
	IncomeTypeCommission  = "COMMISSION"
)

type Zero = struct{}

var ClosedChan = makeClosedChan()

func makeClosedChan() chan Zero {
	ch := make(chan Zero)
	close(ch)
	return ch
}
