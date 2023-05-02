package Routes

import "github.com/gofiber/fiber/v2"

func (r *Router) registerPatternsRoutes(routeGroup fiber.Router) {

	symbolGroup := routeGroup.Group("/pattern")

	symbolGroup.Get("/", r.controller.GetPattern)

	symbolGroup.Get("/rsi", r.controller.GetRSI)
	symbolGroup.Get("/macd", r.controller.GetMACD)

}
