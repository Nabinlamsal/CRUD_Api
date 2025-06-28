package controllers

import (
	"time"

	"github.com/Nabinlamsal/fiber-postgress/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// this is the secret key used to generate and check token
var jwtSecret = []byte("123abc")

// register function, user is stored in database after hashing password
func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User)

		// convert json data from body to our struct (username, password)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid Input",
			})
		}

		// hash password so that it's not saved as plain text in database
		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		user.Password = string(hash)

		// create user in database
		if err := db.Create(user).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "User already exists",
			})
		}

		// if everything okay, show this
		return c.JSON(fiber.Map{"message": "Registered Successfully"})
	}
}

// login function, checks user and password, and gives back a token
func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		input := new(models.User)

		// parse input like before
		if err := c.BodyParser(input); err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}

		user := new(models.User)

		// find user by username
		if err := db.Where("username = ?", input.Username).First(user).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		// compare input password and hashed password from DB
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Password is not correct!",
			})
		}

		// create token claims (info inside token)
		claims := jwt.MapClaims{
			"username": user.Username,                         // so we know who the token belongs to
			"exp":      time.Now().Add(time.Hour * 24).Unix(), // expires after 24 hours
		}

		// generate token using claims and secret
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString(jwtSecret)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Could not login"})
		}

		// send back token to frontend/user
		return c.JSON(fiber.Map{"token": t})
	}
}
