package main

import (
	"log"
	"server/routes"
	utils "server/utils"
)

// @title			Events MIREA Project
// @version		1.0
// @description	Project by Anton Larionov for Events booking
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
func main() {
	utils.InitDotEnv()
	utils.ConnectDb()
	app := utils.InitializeFiber()

	routes.Router(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error in setting up server!")
	}
}
