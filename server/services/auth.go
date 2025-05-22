package services

import (
	"server/middleware"
	"server/models/dto"
	"server/models/entity"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func FindUserByEmail(email string) (*entity.User, error) {
	database := utils.DB

	var foundUser entity.User

	if isError := database.First(&foundUser, "email = ?", email).Error; isError != nil {
		return nil, isError
	}

	return &foundUser, nil
}

func checkUserPassword(userLoginDto *dto.UserLoginDto) bool {
	foundUser, _ := FindUserByEmail(userLoginDto.Email)

	return middleware.CheckHashPassword(userLoginDto.Password, foundUser.PasswordHash)
}

// Register Account Swagger Doc
//
//	@Summary		Register User
//	@Description	Create User by Email Name, Password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			UserData	body		dto.UserRegisterDto	false	"User email, password, role and full name"
//	@Success		201			{string}	string
//	@Failure		400			{object}	models.ApiErrorWrapper
//	@Router			/auth/register [post]
func Register(c *fiber.Ctx) error {
	database := utils.DB

	userDto := dto.UserRegisterDto{}

	if err := c.BodyParser(&userDto); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "ok": false})
	}

	hashedPassword, errorHash := middleware.HashPassword((&userDto).PasswordHash)
	if errorHash != nil {
		return c.Status(400).JSON(fiber.Map{"message": "error while creating password", "ok": false})
	}

	userDto.PasswordHash = hashedPassword

	if _, userError := FindUserByEmail((&userDto).Email); userError == nil {

		return c.Status(400).JSON(fiber.Map{"message": "User already registered", "ok": true})
	}

	database.Create(&entity.User{
		UserRegisterDto: userDto,
	})

	return c.Status(201).SendString("User successfully registered!")
}

// Login user with jwt token
//
//	@Summary		Login User
//	@Description	Login with email and receive JWT Token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			UserCred	body		dto.UserLoginDto	false	"User email and password"
//	@Success		200			{object}	dto.UserResponseDto
//	@Failure		400			{object}	models.ApiErrorWrapper
//	@Failure		404			{object}	models.ApiErrorWrapper
//	@Failure		500			{object}	models.ApiErrorWrapper
//	@Router			/auth/login [post]
func Login(c *fiber.Ctx) error {
	userLogin := dto.UserLoginDto{}

	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "ok": false})
	}

	if _, userError := FindUserByEmail((&userLogin).Email); userError != nil {
		return c.Status(404).JSON(fiber.Map{"message": "This user is not registered", "ok": true})
	}

	if isCorrectPassword := checkUserPassword(&userLogin); !isCorrectPassword {
		return c.Status(400).JSON(fiber.Map{"message": "Wrong Password!", "ok": true})
	}

	token, tokenErr := middleware.CreateToken(userLogin)

	if tokenErr != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Token generation error", "ok": false})
	}

	return c.Status(200).JSON(token)
}
