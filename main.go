package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/quarterblue/parsec/blockchain"
)

func main() {

	app := fiber.New()

	bchain := blockchain.CreateBlockchain()
	bchain.AddBlock("Hello World")
	bchain.AddBlock("Hello Again World!")
	bchain.AddBlock("Hello Again & Again World")
	bchain.AddBlock("Hello Again & Again & Again World")
	bchain.PrintChain()

	app.Get("/api/blockchain", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(bchain)
	})

	app.Post("/api/mine", func(c *fiber.Ctx) error {
		type Request struct {
			Title string `json:"title"`
		}
		var body Request

		err := c.BodyParser(&body)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse json",
			})
		}

		bchain.AddBlock(body.Title)
		return c.Status(fiber.StatusCreated).JSON(bchain)
	})

	log.Fatal(app.Listen(":3000"))

}
