package enum

type OrderType int

const (
	Limit BuyOrSell = iota
	Market
)
