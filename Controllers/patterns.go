package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"trading/Logic/Talib"
)

func (controller *Controller) getPattern(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	interval := c.Params("interval")

	candleSticks, err := controller.marketData.GetCandlesticks(symbol, interval)
	if err != nil {
		return err
	}

	tq := Talib.TaQuote{}
	tq.Init(candleSticks, symbol)

}
