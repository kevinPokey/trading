package MarketData

import "trading/Models"

type MarketData interface {
	Init(apiKey string, secretKey string) (err error)

	GetSymbolPrice(symbol string) (price string, err error)
	GetSymbols() (symbols []string, err error)

	GetCandlesticks(symbol string, interval string) (candleSticks []Models.CandleStick, err error)
}
