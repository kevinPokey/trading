package Controllers

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) GetSymbolPrice(c *fiber.Ctx) error {
	symbol := c.Params("symbol")

	if err := controller.checkTicker(symbol); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	price, err := controller.marketData.GetSymbolPrice(symbol)
	if err != nil {
		return err
	}

	return c.JSON(price)
}

func (controller *Controller) GetSymbols(c *fiber.Ctx) error {

	symbols, err := controller.marketData.GetSymbols()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(symbols)
}
