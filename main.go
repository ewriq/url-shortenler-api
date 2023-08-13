package main

import (
	"fmt"
	"log"
	"url-api/Database"
	"url-api/Utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Link struct {
	Url string `json:"url"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.Config{
		Format: "${ip}:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/:tokens", func(c *fiber.Ctx) error {
		Token :=c.Params("tokens")
		Find, err := Database.Find(Token)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(Find)
		return c.JSON(fiber.Map{
			"data": Find,
			"status": "OK",
		})
	})

	app.Post("/addlink", func(c *fiber.Ctx) error {
		var link Link
		if err := c.BodyParser(&link); err != nil {
			fmt.Println("Error parsing request:", err)
			return err
		}

		url := link.Url
		T0oken := Utils.Token(4)
		Database.Insert(url, T0oken)

		return c.JSON(fiber.Map{
			"token": T0oken,
			"url":   url,
			"msg":   "Başarıyla eklendi!",
			"short": "http://localhost:3000/" + T0oken,
		})
	})
	app.Listen(":3000")
}
