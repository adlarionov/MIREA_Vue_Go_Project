package utils

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func InitializeFiberLogger(c *fiber.Ctx, logString []byte) {

	dirErr := os.Mkdir("./logs", os.ModeAppend)

	if dirErr != nil {
		if !os.IsExist(dirErr) {
			panic(dirErr)
		}
	}

	logFile, logErr := os.OpenFile("./logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if logErr != nil {
		c.Status(500).JSON(fiber.Map{"message": "Something went wrong while adding data to Logs", "ok": false})
		panic(logErr)
	}
	defer logFile.Close()

	logFile.Write(logString)
}
