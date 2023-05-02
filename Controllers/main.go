package Controllers

import (
	"errors"
	"trading/MarketData"
)

type Controller struct {
	marketData MarketData.MarketData

	Constants
}

type Constants struct {
	validTickers    []string
	validIntervals  []string
	validIndicators []string
}

func NewController(marketData MarketData.MarketData, constIndicators, constIntervals, constTickers []string) *Controller {

	return &Controller{
		marketData,
		Constants{constTickers, constIntervals, constIndicators},
	}
}

func isInArray(str string, arr []string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func (controller *Controller) checkTicker(ticker string) error {
	if !isInArray(ticker, controller.validTickers) {
		return errors.New("ticker symbol \"" + ticker + "\" is not valid")
	}
	return nil
}

func (controller *Controller) checkInterval(interval string) error {
	if !isInArray(interval, controller.validIntervals) {
		return errors.New("interval timeframe \"" + interval + "\" is not valid")
	}
	return nil
}

func (controller *Controller) checkIndicators(indicator string) error {
	if !isInArray(indicator, controller.validIndicators) {
		return errors.New("indicator \"" + indicator + "\" is not valid")
	}
	return nil
}
