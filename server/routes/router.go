package routes

import (
	"server/constants"
	"server/middleware"
	"server/services"
	"server/utils"

	_ "server/docs"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func Router(app *fiber.App) {
	// Logging To Log File
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeZone:   "Russia/Moscow",
		Done:       utils.InitializeFiberLogger,
		TimeFormat: "01.01.1970",
	}))

	auth := app.Group("/auth")
	auth.Post("/register", services.Register)
	auth.Post("/login", services.Login)

	app.Get("/swagger/*", swagger.HandlerDefault)

	protected := app.Group("")
	protected.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(constants.Env.APP_JWT_SECRET)},
		ContextKey: constants.JWTContextKey,
	}))
	protected.Use(middleware.InterceptTokenError)
	// Events
	protected.Get("/events", services.GetEvents)
	protected.Get("/events/:id", services.GetEventById)
	protected.Post("/events", services.CreateEvent)
	protected.Put("/events", services.EditEvent)
	protected.Delete("/events/:id", services.DeleteEventById)
}
