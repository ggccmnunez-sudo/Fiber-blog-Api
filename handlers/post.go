package handlers

import (
	"fiber-blog-api/data"
	"fiber-blog-api/dto"
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

	//calling the middleware
	userID := c.Locals("user_id").(uint)
	post.UserID = userID

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
	id := c.Params("id")

	if err := data.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "post not found",
		})
	}

	//calling the middleware
	userId := c.Locals("user_id").(uint)
	if post.UserID != userId {
		return c.Status(403).JSON(fiber.Map{
			"error": "You can only update your own post",
		})
	}

	//creating post
	var inputs dto.CreatePost
	if err := c.Bind().Body(&inputs); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Post doesn't exist",
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

	userId := c.Locals("user_id").(uint)
	if post.UserID != userId {
		return c.Status(403).JSON(fiber.Map{
			"error": "You can only delete your own post",
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
