package main

import (
	"fiber-blog-api/data"
	"fiber-blog-api/models"
	"fiber-blog-api/router"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	data.ConnectDatabase()

	err := data.DB.AutoMigrate(&models.Post{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	app := fiber.New()

	router.BlogRouter(app)

	log.Fatal(app.Listen(":3000"))
}
