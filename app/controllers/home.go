package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"HeadTitle": "Welcome to the public page",
	})
}
