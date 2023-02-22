package controllers

import "github.com/gofiber/fiber/v2"

func GetReferral(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{
		"csrfToken":        c.Locals("token"),
		"ReferralUsername": c.Params("username"),
	})
}
