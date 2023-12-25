package main

import (
	"fmt"
	"log"

	"github.com/aldisypu/go-todolist/config"
	"github.com/aldisypu/go-todolist/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	viperConfig := config.NewViper()
	config.NewDatabase(viperConfig)

	app := config.NewFiber(viperConfig)
	routes.SetupRoute(app)
	app.Use(cors.New())

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
