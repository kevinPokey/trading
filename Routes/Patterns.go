package Routes

import "github.com/gofiber/fiber/v2"

func (r *Router) registerPatternsRoutes(routeGroup fiber.Router) {

	symbolGroup := routeGroup.Group("/pattern")

	symbolGroup.Get("/:symbol/:interval", r.controller.GetPattern)

}
