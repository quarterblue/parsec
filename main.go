package main

import (
	"github.com/quarterblue/parsec/blockchain"
)

func main() {

	// app := fiber.New()

	bchain := blockchain.CreateBlockchain()
	bchain.AddBlock("Hello World")
	bchain.AddBlock("Hello Again World!")
	bchain.AddBlock("Hello Again & Again World")
	bchain.PrintChain()

	// app.Get("/api/blockchain", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(bchain)
	// })

	// log.Fatal(app.Listen(":3000"))

}
