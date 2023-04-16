package Controllers

import "trading/MarketData"

type Controller struct {
	marketData MarketData.MarketData
}

func NewController(marketData MarketData.MarketData) *Controller {
	return &Controller{
		marketData: marketData,
	}
}
