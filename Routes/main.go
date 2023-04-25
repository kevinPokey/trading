package Routes

import (
	"github.com/gofiber/fiber/v2"
	"trading/Controllers"
)

type Router struct {
	app        *fiber.App
	controller *Controllers.Controller
}

func CreateRouter(app *fiber.App, controller *Controllers.Controller) *Router {
	return &Router{
		app:        app,
		controller: controller,
	}
}

func (r *Router) RegisterRoutes() {
	api := r.app.Group("/api")

	r.registerSymbolRoutes(api)
	r.registerCandleStickRoutes(api)
	r.registerPatternsRoutes(api)
}
