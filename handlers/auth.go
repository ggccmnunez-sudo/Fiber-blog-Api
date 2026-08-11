package handlers

import (
	"fiber-blog-api/data"
	"fiber-blog-api/models"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func Hashedpassword(password string) (string, error) {
	//created a variable to use to hashed the password
	//bcrypt.DefaultCost - swap sa 14

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func Register(c fiber.Ctx) error {
	var user models.User
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// call the hashed variable to hash the password
	hashed, err := Hashedpassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to hashed password",
		})
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}
	//Good practice to validate in here after registering if the input is valid like an empty string(""), so that it will not register even if its empty string

	user.Password = hashed

	results := data.DB.Create(&user)
	if results.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Registered successfully",
		"user":    user,
	})
}

func Login(c fiber.Ctx) error {

	var input models.LogIn

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	var user models.User
	results := data.DB.Where("email = ?", input.Email).First(&user)
	if results.Error != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if input.Email == "" || input.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to log in",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		fmt.Println("bcrypt compare error:", err) //debugger

		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	fmt.Println("Found user:", user.Email, user.Password) //debugger

	fmt.Println("Input password:", input.Password) //debugger

	return c.Status(200).JSON(fiber.Map{
		"message": "Log in successfully",
		"user":    user,
	})

}

// Hard coded
// func TempDelete(c fiber.Ctx) error {
// 	var delete models.User
// 	result := data.DB.Where("email = ?", "@loha12").Delete(&models.User{})

// 	if result.Error != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Failed to delete user",
// 		})
// 	}
// 	return c.Status(200).JSON(fiber.Map{
// 		"message": "User deleted successfully",
// 		"delete":  delete,
// 	})
// }
