package exchange_api

import "my.exbot/internal/dto"

type ExchangeApi interface {
	getAssets() dto.AssetDtos
	getAsset(symbol string) dto.AssetDto
	placeOrder(input dto.PlaceOrderDto) dto.PlaceOrderRespDto
	getKline(input dto.GetKlineReqDto) dto.GetKlineRespDto
}
