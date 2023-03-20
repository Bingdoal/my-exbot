package dto

import (
	"github.com/shopspring/decimal"
	"my.exbot/internal/enum"
)

type PlaceOrderDto struct {
	Base      string
	Target    string
	BuyOrSell enum.BuyOrSell
	Type      enum.OrderType
	Price     decimal.Decimal
	Number    decimal.Decimal
}

type PlaceOrderRespDto struct {
	Success bool
	Error   error
}
