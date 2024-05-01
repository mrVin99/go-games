package main

import (
	"games/blocks"
	"games/pkg/cache"
	"games/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
)

var (
	app  = fiber.New()
	db   = database.Conn()
	memo = cache.Conn()
)

func main() {
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​",
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	app.Get("/ws/blocks", blocks.Play(memo, db))
	app.Get("/ui", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})

	log.Println("Server Listening on :8080")
	log.Fatal(app.Listen(":8080"))
}
