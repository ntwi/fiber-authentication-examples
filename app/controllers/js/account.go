package js

import (
	"github.com/gofiber/fiber/v2"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login_js", fiber.Map{})
}

func GetSignUp(c *fiber.Ctx) error {
	return c.Render("signup_js", fiber.Map{})
}

func GetAccount(c *fiber.Ctx) error {
	return c.Render("account_js", fiber.Map{})
}
