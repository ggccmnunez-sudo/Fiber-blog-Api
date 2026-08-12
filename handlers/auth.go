package handlers

import (
	"fiber-blog-api/data"
	"fiber-blog-api/dto"
	"fiber-blog-api/models"
	"fiber-blog-api/utils"

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
	var register dto.RegisterInputs
	if err := c.Bind().Body(&register); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	//Good practice to validate in here after registering if the input is valid like an empty string("")
	// so that it will not register even if its an empty string
	if register.Email == "" || register.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}
	// call the hashed variable to hash the password
	hashed, err := Hashedpassword(register.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to hashed password",
		})
	}

	register.Password = hashed

	user := models.User{
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password,
	}
	results := data.DB.Create(&user)

	if results.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	token, err := utils.GenerateJwt(user.ID, user.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Registered successfully",
		"user":    user,
		"token":   token,
	})
}

func Login(c fiber.Ctx) error {

	var loginInput dto.LoginInput

	if err := c.Bind().Body(&loginInput); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if loginInput.Email == "" || loginInput.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to log in",
		})
	}

	var user models.User
	results := data.DB.Where("email = ?", loginInput.Email).First(&user)
	if results.Error != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {

		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	token, err := utils.GenerateJwt(user.ID, user.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Log in successfully",
		"user":    user,
		"token":   token,
	})

}
