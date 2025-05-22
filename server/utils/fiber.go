package utils

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeFiber() *fiber.App {
	app := fiber.New()

	return app
}
