package router

import (
	"auth/app/controllers"
	"auth/app/controllers/js"
	"auth/app/middlewares/anonym_auth"
	"auth/app/middlewares/cookie_auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"os"
	"time"
)

type WebRouter struct {
}

func (h WebRouter) InstallRouter(app *fiber.App) {
	anonymMiddleware := anonym_auth.New(anonym_auth.Config{
		RedirectUrl: "/account",
		HeaderName:  "Authorization",
		CookieName:  "auth",
	})
	cookieAuthMiddleware := cookie_auth.New(cookie_auth.Config{
		Secret:      os.Getenv("JWT_SECRET"),
		CookieName:  "auth",
		RedirectUrl: "/login",
	})
	csrfMiddleware := csrf.New(csrf.Config{
		KeyLookup:      "form:_csrf",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ContextKey:     "token",
	})
	group := app.Group("", cors.New(), csrfMiddleware)
	group.Get("/", controllers.GetHome)
	group.Get("/login", anonymMiddleware, controllers.GetLogin)
	group.Post("/login", anonymMiddleware, controllers.PostLogin)
	group.Get("/signup", anonymMiddleware, controllers.GetSignUp)
	group.Post("/signup", anonymMiddleware, controllers.PostSignUp)
	group.Get("/account", cookieAuthMiddleware, controllers.GetAccount)
	group.Get("/r/:username", controllers.GetReferral)
	group.Post("/logout", cookieAuthMiddleware, controllers.PostLogout)

	jsGroup := app.Group("/js", cors.New())
	jsGroup.Get("/login", js.GetLogin)
	jsGroup.Get("/signup", js.GetSignUp)
	jsGroup.Get("/account", js.GetAccount)
	jsGroup.Get("/r/:username", js.GetReferral)
}

func NewWebRouter() *WebRouter {
	return &WebRouter{}
}
