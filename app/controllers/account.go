package controllers

import (
	"auth/app/models"
	"auth/app/services"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"csrfToken": c.Locals("token"),
	})
}

func PostLogin(c *fiber.Ctx) error {
	loginModel := &models.LoginModel{}
	err := c.BodyParser(loginModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("login", fiber.Map{
			"csrfToken": c.Locals("token"),
			"error":     "Invalid login credentials",
		})
	}

	token, err := services.Login(loginModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("login", fiber.Map{
			"csrfToken": c.Locals("token"),
			"error":     "Invalid login credentials",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    url.PathEscape(token),
		HTTPOnly: true,
		SameSite: "strict",
		Secure:   true,
	})
	return c.Redirect("/account")
}

func GetSignUp(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{
		"csrfToken": c.Locals("token"),
	})
}

func PostSignUp(c *fiber.Ctx) error {
	signUpModel := &models.SignUpModel{}
	err := c.BodyParser(signUpModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("login", fiber.Map{
			"csrfToken": c.Locals("token"),
			"error":     "Invalid login credentials",
		})
	}

	err = services.SignUp(signUpModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("register", fiber.Map{
			"csrfToken": c.Locals("token"),
			"error":     err,
		})
	}

	return c.Redirect("/login")
}

func GetAccount(c *fiber.Ctx) error {
	username := c.Locals("username")
	return c.Render("account", fiber.Map{
		"csrfToken": c.Locals("token"),
		"username":  username,
	})
}

func PostLogout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    "",
		HTTPOnly: true,
		SameSite: "strict",
		Secure:   true,
	})
	return c.Redirect("/")
}
