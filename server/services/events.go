package services

import (
	"server/models/dto"
	"server/models/entity"
	"server/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Get All Events
//
//	@Summary		Get All Events
//	@Description	Get All Events without ID
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Success		200				{object}	[]entity.Event
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Router			/events [get]
func GetEvents(c *fiber.Ctx) error {
	database := utils.DB

	events := []entity.Event{}

	if database.Find(&events).Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error while finding events in DB", "ok": false})
	}

	return c.Status(200).JSON(events)
}

// Get Event By Id
//
//	@Summary		Get Event
//	@Description	Get Event by event id
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Param			id				path		int		true	"Event ID"
//	@Success		200				{object}	entity.Event
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/events/{id} [get]
func GetEventById(c *fiber.Ctx) error {
	database := utils.DB

	event := entity.Event{}

	idParam := c.Params("id")

	eventId, convError := strconv.Atoi(idParam)

	if convError != nil {
		return c.Status(500).JSON(fiber.Map{"message": "", "ok": false})
	}

	if database.Find(&event, eventId).Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error while finding events in DB", "ok": false})
	}

	if eventId != int(event.EventId) {
		return c.Status(404).JSON(fiber.Map{"message": "Event not found", "ok": true})
	}

	return c.Status(200).JSON(event)
}

// Create Event
//
//	@Summary		Create Event
//	@Description	Create Event by body params
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Param			Event			body		dto.EventRequestDto	false	"Event Data"
//	@Success		201				{string}	string
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/events [post]
func CreateEvent(c *fiber.Ctx) error {
	database := utils.DB

	event := dto.EventRequestDto{}

	if bodyError := c.BodyParser(&event); bodyError != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Wrong body parameters", "ok": false})
	}

	user, userErr := FindUserByEmail(event.OwnerEmail)

	if userErr != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found", "ok": true})
	}

	newEvent := entity.Event{
		OwnerId:   user.UserId,
		EventDto:  event.Event,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if database.Create(&newEvent).Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong while saving event to DB", "ok": false})
	}

	return c.Status(201).SendString("Event has been succesfully created!")
}

// Edit Event
//
//	@Summary		Edit Event
//	@Description	Edit Event by body params
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Param			Event			body		dto.EventRequestDto	false	"Event Edit Data"
//	@Success		200				{string}	string
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/events [put]
func EditEvent(c *fiber.Ctx) error {
	database := utils.DB

	event := dto.EventRequestDto{}

	if c.BodyParser(&event) != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Wrong body parameters", "ok": false})
	}

	eventToEdit := entity.Event{}

	if database.First(&eventToEdit, "Title = ?", event.Event.Title).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "No such event in DB", "ok": true})
	}

	if database.Save(&entity.Event{
		EventId:  eventToEdit.EventId,
		EventDto: event.Event,
	}).Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error while updating event in DB", "ok": false})
	}

	return c.Status(200).SendString("Event successfully updated!")
}

// Delete Event By Id
//
//	@Summary		Delete Event
//	@Description	Delete Event by id
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Param			id				path		int		true	"Event Id"
//	@Success		200				{string}	string
//	@Failure		500				{object}	models.ApiErrorWrapper
//	@Failure		404				{object}	models.ApiErrorWrapper
//	@Router			/events/{id} [delete]
func DeleteEventById(c *fiber.Ctx) error {
	database := utils.DB

	idParam := c.Params("id")
	eventId, convError := strconv.Atoi(idParam)

	if convError != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Id param parse error", "ok": false})
	}

	if database.Delete(&entity.Event{}, eventId).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "No such event or event has already been deleted", "ok": true})
	}

	return c.Status(200).SendString("Event successfully deleted!")
}
