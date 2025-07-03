package main

import (
	"log"

	"github.com/PirateWindows/DriverMockup/internal/router"
)

func main() {
	r := router.SetupRouter()
	log.Println("Starting server on :8080")
	r.Run(":8080") // Using Gin

}
