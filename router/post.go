package router

import (
	"fiber-blog-api/handlers"
	"fiber-blog-api/middleware"

	"github.com/gofiber/fiber/v3"
)

func BlogRouter(app *fiber.App) {
	app.Get("/post", handlers.GetPost)
	app.Get("/post/:id", handlers.GetPostById)
	app.Post("/post", middleware.Protected(), handlers.CreatePost)
	app.Put("/post/:id", middleware.Protected(), handlers.UpdatePost)
	app.Delete("/post/:id", middleware.Protected(), handlers.DeletePost)

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

}
