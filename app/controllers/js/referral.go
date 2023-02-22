package js

import "github.com/gofiber/fiber/v2"

func GetReferral(c *fiber.Ctx) error {
	return c.Render("signup_js", fiber.Map{})
}
