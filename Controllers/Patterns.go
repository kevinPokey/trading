package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"trading/Logic/Talib"
)

func (controller *Controller) GetPattern(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	interval := c.Params("interval")

	candleSticks, err := controller.marketData.GetCandlesticks(symbol, interval)
	if err != nil {
		return err
	}

	tq := Talib.TaQuote{}
	tq.Init(candleSticks, symbol)

	result := tq.GetRSI(14)
	return c.JSON(result)
}
