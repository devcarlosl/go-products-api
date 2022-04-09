package main

import (
	"github.com/devcarlosl/go-products-api/pkg/common/config"
	"github.com/devcarlosl/go-products-api/pkg/common/db"
	"github.com/devcarlosl/go-products-api/pkg/products"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	products.RegisterRoutes(app, h)

	app.Listen(c.Port)
}