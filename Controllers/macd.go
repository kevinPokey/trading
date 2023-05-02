package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"trading/Logic/Talib"
	"trading/Models"
)

func (controller *Controller) GetMACD(c *fiber.Ctx) error {
	si := new(Models.SymbolInterval)

	if err := c.BodyParser(si); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := controller.checkTicker(si.Symbol); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := controller.checkInterval(si.Interval); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	candleSticks, err := controller.marketData.GetCandlesticks(si.Symbol, si.Interval)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	tq := Talib.TaQuote{}
	tq.Init(candleSticks, si.Symbol)
	macd, signal, bars := tq.GetMACD(12, 26, 9)

	result := Models.Macd{Macd: macd, Signal: signal, Bars: bars}

	return c.JSON(result)
}
