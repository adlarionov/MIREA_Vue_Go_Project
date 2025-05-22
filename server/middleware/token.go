package middleware

import (
	"server/constants"
	"server/models/dto"
	"server/models/entity"
	"server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user dto.UserLoginDto) (dto.UserResponseDto, error) {
	encryptedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})

	token, tokenErr := encryptedToken.SignedString([]byte(constants.Env.APP_JWT_SECRET))

	return dto.UserResponseDto{
		Token: token,
	}, tokenErr
}

func CheckToken(c *fiber.Ctx) (jwt.MapClaims, bool) {
	jwtToken, ok := c.Context().Value(constants.JWTContextKey).(*jwt.Token)

	if !ok {
		return nil, false
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, false
	}

	return payload, true
}

func InterceptTokenError(c *fiber.Ctx) error {
	database := utils.DB
	jwtPayload, ok := CheckToken(c)

	if !ok {
		return c.Status(500).JSON(fiber.Map{"message": "Error while taking user token", "ok": false})
	}

	user := entity.User{}

	if database.First(&user, "email = ?", jwtPayload["sub"]).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found", "ok": true})
	}

	return c.Next()
}
