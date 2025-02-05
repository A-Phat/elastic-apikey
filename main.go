package main

import (
	"elastic-apikey/config"
	"elastic-apikey/routes"
	"log"
)

func main() {
	config.InitElastic()

	r := routes.SetupRouter()
	log.Println("Server is running on port 8080")
	r.Run(":8080")
}
