package services

import (
	"fmt"
	"server/models/dto"
	"server/models/entity"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

// Get All Venues
//
//	@Summary		Get All Venues
//	@Description	Get All venues without ID
//	@Tags			Venues
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"	default(Bearer )
//	@Success		200				{object}	[]entity.Venue
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Router			/venues [get]
func GetVenues(c *fiber.Ctx) error {
	database := utils.DB

	venues := []entity.Venue{}

	if createError := database.Model(&entity.Venue{}).Preload("Image").Find(&venues).Error; createError != nil {
		return c.Status(500).JSON(fiber.Map{"message": fmt.Sprintf("Error while finding venues in DB %s", createError), "ok": false})
	}

	return c.Status(200).JSON(venues)
}

// Get Venue By Id
//
//	@Summary		Get Venue
//	@Description	Get Venue by venue id
//	@Tags			Venues
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"	default(Bearer )
//	@Param			id				path		int		true	"Venue ID"		default(1)
//	@Success		200				{object}	entity.Venue
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/venues/{id} [get]
func GetVenueById(c *fiber.Ctx) error {
	database := utils.DB

	venue := entity.Venue{}

	venueId, venueIdError := c.ParamsInt("id")

	if venueIdError != nil {
		return c.Status(500).JSON(fiber.Map{"message": venueIdError, "ok": false})
	}

	if database.Preload("Image").Find(&venue, venueId).Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error while finding event by id in DB", "ok": false})
	}

	return c.Status(200).JSON(venue)
}

// Create Venue
//
//	@Summary		Create Venue
//	@Description	Create Venue by venue id
//	@Tags			Venues
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"	default(Bearer )
//	@Param			Venue			body		dto.VenueRequestDto	false	"Venue Data"
//	@Success		201				{string}	string
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/venues [post]
func CreateVenue(c *fiber.Ctx) error {
	database := utils.DB

	venueBody := &dto.VenueRequestDto{}

	if c.BodyParser(&venueBody) != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Wrong body parameters", "ok": false})
	}

	if createError := database.Create(&entity.Venue{VenueDto: venueBody.Venue}).Error; createError != nil {
		return c.Status(500).JSON(fiber.Map{"message": fmt.Sprintf("Error while creating venue. %s", createError), "ok": false})
	}

	return c.Status(201).SendString("Successfully created venue!")
}

// Add Image To Existing Venue
//
//	@Summary		Add Image to Venue
//	@Description	Add Image to Venue by venue id
//	@Tags			Venues
//	@Accept			mpfd
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"	default(Bearer )
//	@Param			Image			formData	file	true	"Image to Venue"
//	@Param			id				path		int		true	"Venue ID"	default(0)
//	@Success		200				{object}	entity.Image
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/venues/{id}/images [post]
func AddImageToVenueById(c *fiber.Ctx) error {
	database := utils.DB

	venueModel := &entity.Venue{}

	venueId, venueIdError := c.ParamsInt("id")

	if venueIdError != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Wrong venue id format type", "ok": true})
	}

	if database.Find(venueModel, venueId).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Venue not found", "ok": true})
	}

	image, imageErr := c.FormFile("Image")

	if imageErr != nil {
		return c.Status(400).JSON(fiber.Map{"message": imageErr, "ok": true})
	}

	pathname := fmt.Sprintf("./static/%s", image.Filename)

	if saveFileError := c.SaveFile(image, pathname); saveFileError != nil {
		return c.Status(500).JSON(fiber.Map{"message": fmt.Sprintf("Error while saving file. %s", saveFileError), "ok": false})
	}

	imageDtoModel := &dto.ImageDto{ImageUrl: pathname}
	imageModel := &entity.Image{ImageDto: *imageDtoModel}

	if createError := database.Create(imageModel).Error; createError != nil {
		return c.Status(500).JSON(fiber.Map{"message": fmt.Sprintf("Error while saving file. %s", createError), "ok": false})
	}

	if updatedVenueError := database.Model(&entity.Venue{}).Where("id", venueId).Update("image_id", imageModel.ID).Error; updatedVenueError != nil {
		return c.Status(500).JSON(fiber.Map{"message": updatedVenueError, "ok": false})
	}

	return c.Status(200).JSON(imageModel)
}
