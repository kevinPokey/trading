package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"trading/Logic/Talib"
	"trading/Models"
)

func (controller *Controller) GetPattern(c *fiber.Ctx) error {

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
		return err
	}

	tq := Talib.TaQuote{}
	tq.Init(candleSticks, si.Symbol)

	var result []string

	return c.JSON(result)
}
