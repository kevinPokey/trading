package Routes

import "github.com/gofiber/fiber/v2"

func (r *Router) registerCandleStickRoutes(routeGroup fiber.Router) {

	symbolGroup := routeGroup.Group("/candle")

	symbolGroup.Get("/:symbol/:interval", r.controller.GetCandleSticks)

}
