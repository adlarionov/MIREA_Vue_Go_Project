package services

import (
	"server/models/entity"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

// Get All Bookings
//
//	@Summary		Get All Bookings
//	@Description	Get All Bookings without ID
//	@Tags			Bookings
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"	default(Bearer )
//	@Success		200				{object}	[]entity.Booking
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Router			/bookings [get]
func GetBookings(c *fiber.Ctx) error {
	database := utils.DB

	bookings := []entity.Booking{}

	if database.Model(&entity.Booking{}).Preload("Event").Preload("Venue").Preload("Venue.Image").Find(&bookings).Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error while finding bookings in DB", "ok": false})
	}

	return c.Status(200).JSON(bookings)
}

// Get Booking By Id
//
//	@Summary		Get Booking
//	@Description	Get Booking by booking id
//	@Tags			Bookings
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"	default(Bearer )
//	@Param			id				path		int		true	"Booking ID"
//	@Success		200				{object}	entity.Booking
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/bookings/{id} [get]
func GetBookingById(c *fiber.Ctx) error {
	database := utils.DB

	bookingModel := entity.Booking{}

	bookingId, bookingIdError := c.ParamsInt("id")

	if bookingIdError != nil {
		return c.Status(500).JSON(fiber.Map{"message": bookingIdError, "ok": false})
	}

	if database.Model(&bookingModel).Preload("Event").Preload("Venue").Preload("Venue.Image").Find(&bookingModel, bookingId).Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error while finding booking in DB", "ok": false})
	}

	if bookingId != int(bookingModel.ID) {
		return c.Status(404).JSON(fiber.Map{"message": "Booking not found", "ok": true})
	}

	return c.Status(200).JSON(bookingModel)
}

// Create Booking
//
//	@Summary		Create Booking
//	@Description	Create Booking by body params
//	@Tags			Bookings
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer Token"	default(Bearer )
//	@Param			Booking			body		entity.BookingRequestDto	false	"Booking Data"
//	@Success		201				{string}	string
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/bookings [post]
func CreateBooking(c *fiber.Ctx) error {
	database := utils.DB

	bookingModel := entity.BookingRequestDto{}

	if bodyError := c.BodyParser(&bookingModel); bodyError != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Wrong body parameters", "ok": false})
	}

	newBooking := entity.Booking{
		BookingDto: bookingModel.BookingDto,
		EventId:    bookingModel.EventId,
		VenueId:    bookingModel.VenueId,
	}

	if database.Create(&newBooking).Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong while saving booking to DB", "ok": false})
	}

	return c.Status(201).SendString("Booking has been succesfully created!")
}
