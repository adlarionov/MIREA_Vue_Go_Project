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

func CreateToken(loginDto dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	encryptedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": loginDto.Email,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})

	token, tokenErr := encryptedToken.SignedString([]byte(constants.Env.APP_JWT_SECRET))

	return dto.LoginResponseDto{
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
		return c.Status(500).JSON(fiber.Map{"message": "Error while taking organization token", "ok": false})
	}

	organization := entity.Organization{}

	if database.First(&organization, "email = ?", jwtPayload["sub"]).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Organization not found", "ok": true})
	}

	c.Locals(constants.OrganizationEmailKey, jwtPayload["sub"])
	c.Locals(constants.OrganizationIDKey, organization.ID)

	return c.Next()
}
