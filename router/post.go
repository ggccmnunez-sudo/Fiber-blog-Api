package router

import (
	"fiber-blog-api/handlers"

	"github.com/gofiber/fiber/v3"
)

func BlogRouter(app *fiber.App) {
	app.Get("/post", handlers.GetPost)
	app.Get("/post/:id", handlers.GetPostById)
	app.Post("/post", handlers.CreatePost)
	app.Put("/post/:id", handlers.UpdatePost)
	app.Delete("/post/:id", handlers.DeletePost)

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// app.Delete("/delete", handlers.TempDelete)
}
