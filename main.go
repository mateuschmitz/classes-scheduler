package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mateuschmitz/classes-scheduler/bookings"
	"github.com/mateuschmitz/classes-scheduler/classes"
	"github.com/mateuschmitz/classes-scheduler/database"
)

func main() {
    app := fiber.New()
    database.ConnectDB()
    defer database.DB.Close()

    api := app.Group("/api")
    classes.Register(api, database.DB)
    bookings.Register(api, database.DB)

    log.Fatal(app.Listen(":8080"))
}