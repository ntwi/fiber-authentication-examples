package cookie_auth

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type Config struct {
	Filter       func(c *fiber.Ctx) bool
	Unauthorized fiber.Handler
	Decode       func(c *fiber.Ctx) (*jwt.MapClaims, error)
	Secret       string
	Expiry       int64
	CookieName   string
	RedirectUrl  string
}

var ConfigDefault = Config{
	Filter:       nil,
	Decode:       nil,
	Unauthorized: nil,
	Secret:       os.Getenv("JWT_SECRET"),
	Expiry:       60,
	CookieName:   "auth",
	RedirectUrl:  "/login",
}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	if cfg.Secret == "" {
		cfg.Secret = ConfigDefault.Secret
	}

	if cfg.Expiry == 0 {
		cfg.Expiry = ConfigDefault.Expiry
	}

	if cfg.Decode == nil {
		cfg.Decode = func(c *fiber.Ctx) (*jwt.MapClaims, error) {

			authCookie := c.Cookies(cfg.CookieName)
			if authCookie == "" {
				return nil, errors.New("auth cookie_auth is empty")
			}

			token, err := jwt.Parse(
				authCookie,
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(cfg.Secret), nil
				},
			)

			if err != nil {
				return nil, errors.New("error parsing token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !(ok && token.Valid) {
				return nil, errors.New("invalid token")
			}

			if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
				return nil, errors.New("jwt_auth is expired")
			}

			return &claims, nil
		}
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.Redirect(cfg.RedirectUrl)
		}
	}

	return cfg
}

func GenerateToken(claims *jwt.MapClaims, expiryAfter int64) (string, error) {
	if expiryAfter == 0 {
		expiryAfter = ConfigDefault.Expiry
	}

	(*claims)["exp"] = time.Now().UTC().Unix() + expiryAfter

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", errors.New("error creating a token")
	}

	return signedToken, nil
}

func New(config Config) fiber.Handler {
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		claims, err := cfg.Decode(c)

		if err == nil {
			c.Locals("username", (*claims)["username"])
			return c.Next()
		}

		c.ClearCookie(cfg.CookieName)
		return cfg.Unauthorized(c)
	}
}
