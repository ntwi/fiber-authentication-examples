package router

import (
	endpoints "auth/app/controllers/api"
	"auth/app/middlewares/jwt_auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())
	accountApi := api.Group("/account")
	accountApi.Get("/", jwt_auth.Protected(), endpoints.GetAccount)
	accountApi.Post("/login", endpoints.PostLogin)
	accountApi.Post("/signup", endpoints.PostSignUp)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
