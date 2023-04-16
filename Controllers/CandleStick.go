package Controllers

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) GetCandleSticks(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	interval := c.Params("interval")

	candleSticks, err := controller.marketData.GetCandlesticks(symbol, interval)
	if err != nil {
		return err
	}

	return c.JSON(candleSticks)
}
