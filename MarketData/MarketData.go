package MarketData

import "awesomeProject/Models"

type MarketData interface {
	Init() (err error)

	GetSymbolPrice(symbol string) (price string, err error)
	GetKlines(symbol string, interval int) (candleSticks Models.CandleStick, err error)
}
