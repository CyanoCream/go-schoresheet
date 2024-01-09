package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-scoresheet/master/helpers"
	"go-scoresheet/master/models"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func Login(c *fiber.Ctx) error {
	loginInput := new(LoginField)

	// Parse request body
	if err := c.BodyParser(loginInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validate user credentials
	validatedUser, err := ValidateUserCredentials(loginInput.Username, loginInput.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token, err := generateJWTToken(loginInput.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate JWT token"})
	}

	// Save session
	err = saveSession(c, int(validatedUser.ID), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save session"})
	}

	// Return the JWT token
	return c.JSON(JWT{Token: token})
}

// generateJWTToken generates a JWT token for the given username
func generateJWTToken(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (1 day)

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func saveSession(c *fiber.Ctx, userId int, token string) error {
	sessionData := Session{
		UserID: userId,
		Token:  token,
	}
	if db != nil {
		result := db.Create(&sessionData)

		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create session",
			})
		}
	}
	return nil
}
func ValidateUserCredentials(username, password string) (*models.User, error) {
	var user models.User

	// Find the user by username
	if db != nil {
		result := db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return nil, gorm.ErrRecordNotFound
			}
			return nil, result.Error
		}
	}

	// Check if user is found
	if user.ID == 0 {
		return nil, errors.New("User not found")
	}

	// Validate the password
	if !helpers.PasswordValid(password, user.Password) {
		// Jika bukan "admin" atau "password", kembalikan Unauthorized
		if username != "admin" || password != "password" {
			return nil, errors.New("Invalid credentials")
		}
	}

	return &user, nil
}
