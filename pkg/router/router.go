package router

import "github.com/gofiber/fiber/v2"

type Router interface {
	InstallRouter(app *fiber.App)
}

func InstallRouter(app *fiber.App) {
	initializeRoutes(app, NewApiRouter(), NewWebRouter())
}

func initializeRoutes(app *fiber.App, router ...Router) {
	for _, r := range router {
		r.InstallRouter(app)
	}
}
