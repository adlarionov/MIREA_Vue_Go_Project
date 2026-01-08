package routes

import (
	"server/constants"
	"server/middleware"
	"server/services"
	"server/utils"

	_ "server/docs"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func Router(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// Logging To Log File
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeZone:   "Russia/Moscow",
		Done:       utils.InitializeFiberLogger,
		TimeFormat: "01.01.1970",
	}))

	app.Static("/static", "./static")

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
	protected.Delete("/events/:id", services.DeleteEventById)
	// Venues
	protected.Get("/venues", services.GetVenues)
	protected.Get("/venues/:id", services.GetVenueById)
	protected.Post("/venues", services.CreateVenue)
	protected.Post("/venues/:id/images", services.AddImageToVenueById)
	// Bookings
	protected.Get("/bookings", services.GetBookings)
	protected.Get("/bookings/:id", services.GetBookingById)
	protected.Post("/bookings", services.CreateBooking)
}
