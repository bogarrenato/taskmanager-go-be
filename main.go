package main

import (
	"fmt"
	"os"
	"taskmanagerapp/database"
	"taskmanagerapp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	mydir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(mydir)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	app.Listen(":8000")
}
