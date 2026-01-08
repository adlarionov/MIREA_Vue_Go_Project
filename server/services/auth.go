package services

import (
	"server/middleware"
	"server/models/dto"
	"server/models/entity"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func FindOrganizationByEmail(email string) (*entity.Organization, error) {
	database := utils.DB

	var foundOrganization entity.Organization

	if isError := database.First(&foundOrganization, "email = ?", email).Error; isError != nil {
		return nil, isError
	}

	return &foundOrganization, nil
}

func checkPassword(loginRequestDto *dto.LoginRequestDto) bool {
	foundOrganization, _ := FindOrganizationByEmail(loginRequestDto.Email)

	return middleware.CheckHashPassword(loginRequestDto.Password, foundOrganization.PasswordHash)
}

// Register Account Swagger Doc
//
//	@Summary		Register Organization
//	@Description	Create Organization by Email, Password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			OrganizationData	body		dto.RegisterRequestDto	false	"Organization email, password, name of company"
//	@Success		201					{string}	string
//	@Failure		400					{object}	models.ApiErrorWrapper
//	@Router			/auth/register [post]
func Register(c *fiber.Ctx) error {
	database := utils.DB

	authDto := dto.RegisterRequestDto{}

	if err := c.BodyParser(&authDto); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "ok": false})
	}

	hashedPassword, errorHash := middleware.HashPassword((&authDto).PasswordHash)

	if errorHash != nil {
		return c.Status(400).JSON(fiber.Map{"message": "error while creating password", "ok": false})
	}

	authDto.PasswordHash = hashedPassword

	if _, authError := FindOrganizationByEmail((&authDto).Email); authError == nil {

		return c.Status(400).JSON(fiber.Map{"message": "Organization already registered", "ok": true})
	}

	database.Create(&entity.Organization{
		RegisterRequestDto: authDto,
	})

	return c.Status(201).SendString("Organization successfully registered!")
}

// Login organization with jwt token
//
//	@Summary		Login Organization
//	@Description	Login with email and receive JWT Token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			OrganizationCred	body		dto.LoginRequestDto	false	"Organization email and password"
//	@Success		200					{object}	dto.LoginResponseDto
//	@Failure		400					{object}	models.ApiErrorWrapper
//	@Failure		404					{object}	models.ApiErrorWrapper
//	@Failure		500					{object}	models.ApiErrorWrapper
//	@Router			/auth/login [post]
func Login(c *fiber.Ctx) error {
	login := dto.LoginRequestDto{}

	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "ok": false})
	}

	if _, userError := FindOrganizationByEmail(login.Email); userError != nil {
		return c.Status(404).JSON(fiber.Map{"message": "This user is not registered", "ok": true})
	}

	if !checkPassword(&login) {
		return c.Status(400).JSON(fiber.Map{"message": "Wrong Password!", "ok": true})
	}

	token, tokenErr := middleware.CreateToken(login)

	if tokenErr != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Token generation error", "ok": false})
	}

	return c.Status(200).JSON(token)
}
