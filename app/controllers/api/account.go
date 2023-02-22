package api

import (
	"auth/app/models"
	"auth/app/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func PostLogin(c *fiber.Ctx) error {
	loginModel := &models.LoginModel{}
	err := c.BodyParser(loginModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid login credentials",
		})
	}

	token, err := services.Login(loginModel)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid login credentials",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
	})
}

func PostSignUp(c *fiber.Ctx) error {
	signUpModel := &models.SignUpModel{}
	err := c.BodyParser(signUpModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid login credentials",
		})
	}

	err = services.SignUp(signUpModel)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid login credentials",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
}

func GetAccount(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Welcome %s!", claims["username"]),
	})
}
