package Controllers

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) GetSymbolPrice(c *fiber.Ctx) error {
	symbol := c.Params("symbol")

	price, err := controller.marketData.GetSymbolPrice(symbol)
	if err != nil {
		return err
	}

	return c.JSON(price)
}
