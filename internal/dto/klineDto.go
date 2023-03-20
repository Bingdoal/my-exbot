package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type GetKlineReqDto struct {
	Base   string
	Target string
	Limit  int64
}

type GetKlineRespDto struct {
	Open     decimal.Decimal
	Close    decimal.Decimal
	Min      decimal.Decimal
	Max      decimal.Decimal
	Current  decimal.Decimal
	DateTime time.Time
}
