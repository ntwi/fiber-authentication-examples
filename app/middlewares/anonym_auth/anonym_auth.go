package anonym_auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Filter          func(c *fiber.Ctx) bool
	Authorized      fiber.Handler
	CheckAuthorized func(c *fiber.Ctx) error
	CookieName      string
	HeaderName      string
	RedirectUrl     string
}

var ConfigDefault = Config{
	Filter:          nil,
	CheckAuthorized: nil,
	HeaderName:      "Authorization",
	CookieName:      "auth",
	RedirectUrl:     "/login",
}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	if cfg.CheckAuthorized == nil {
		cfg.CheckAuthorized = func(c *fiber.Ctx) error {

			authCookie := c.Cookies(cfg.CookieName)
			if authCookie != "" {
				return errors.New("cookie auth is not empty")
			}

			authHeader := c.Get(cfg.HeaderName)
			if authHeader != "" {
				return errors.New("header auth is not empty")
			}

			return nil
		}
	}

	if cfg.Authorized == nil {
		cfg.Authorized = func(c *fiber.Ctx) error {
			return c.Redirect(cfg.RedirectUrl)
		}
	}

	return cfg
}

func New(config Config) fiber.Handler {
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		err := cfg.CheckAuthorized(c)
		if err == nil {
			return c.Next()
		}

		return cfg.Authorized(c)
	}
}
