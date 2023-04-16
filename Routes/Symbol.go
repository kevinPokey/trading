package Routes

import "github.com/gofiber/fiber/v2"

func (r *Router) registerSymbolRoutes(routeGroup fiber.Router) {

	symbolGroup := routeGroup.Group("/symbol")

	symbolGroup.Get("/:symbol", r.controller.GetSymbolPrice)

}
