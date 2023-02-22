package jwt_auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"

	jwt "github.com/gofiber/jwt/v2"
)

func Protected() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler:  jwtError,
		SigningMethod: "HS512",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	fmt.Println(err)
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"status": "error", "data": nil})
	}

	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "data": nil})
}
