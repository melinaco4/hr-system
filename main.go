package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const dbName = "hr-system"
const mongoURI = "mongodb://localhost:27017" + dbName

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {

}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {

	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
