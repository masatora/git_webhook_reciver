package main

import (
	"git_webhook/api"
	"git_webhook/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Set_env()
	web := fiber.New()
	web.Post("/webhook", api.Webhook)

	log.Fatal(web.Listen(os.Getenv("GO_HOST") + ":" + os.Getenv("GO_PORT")))
}
