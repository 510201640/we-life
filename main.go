package main

import (
	"jaden/we-life/initialize"
)

func main() {

	initialize.SetupConfig()

	initialize.SetUpRedis()

	initialize.SetupDB()

	initialize.SetUpModule()

	initialize.NewLogger()

	initialize.SetupServer()
}
