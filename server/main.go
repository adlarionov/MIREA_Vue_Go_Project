package main

import (
	utils "server/utils"
)

func main() {
	utils.InitDotEnv()
	utils.ConnectDb()
}
