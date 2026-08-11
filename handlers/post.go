package handlers

import (
	"fiber-blog-api/data"
	"fiber-blog-api/models"

	"github.com/gofiber/fiber/v3"
)

func GetPost(c fiber.Ctx) error {
	var post []models.Post

	results := data.DB.Find(&post)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User doesn't exist",
		})
	}

	return c.Status(200).JSON(post)
}

func GetPostById(c fiber.Ctx) error {
	var post models.Post
	id := c.Params("id")

	results := data.DB.First(&post, id)
	if results.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "post not found",
		})
	}

	return c.Status(200).JSON(post)
}

func CreatePost(c fiber.Ctx) error {
	var post models.Post

	if err := c.Bind().Body(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Post doesnt exist",
		})
	}

	results := data.DB.Create(&post)
	if results.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Post added successfully",
		"post":    post,
	})
}

func UpdatePost(c fiber.Ctx) error {
	var post models.Post

	if err := c.Bind().Body(&post); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User doesn't exist",
		})
	}

	results := data.DB.Save(&post)
	if results.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})

	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User updated successfully",
		"post":    post,
	})

}

func DeletePost(c fiber.Ctx) error {
	var post models.Post

	id := c.Params("id")
	if err := data.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	if err := data.DB.Delete(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Post deleted successfully",
		"post":    post,
	})

}
