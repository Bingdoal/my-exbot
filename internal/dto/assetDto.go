package dto

import "github.com/shopspring/decimal"

type AssetDto struct {
	Symbol string
	Number decimal.Decimal
}

type AssetDtos []AssetDto
