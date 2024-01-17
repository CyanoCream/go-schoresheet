package helpers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
)

const secretKey = "t3StS3cR3t!"

func GenerateToken(id uint, email string, roleCodes []string) (token string, err error) {

	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  roleCodes,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = parseToken.SignedString([]byte(secretKey))

	return
}

func VerifyToken(c *fiber.Ctx) error {
	headerToken := c.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer ")

	if !bearer {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Bearer token not found"})
	}

	stringToken := headerToken[7:]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("failed to get sign token")
		}

		// Replace "secretKey" with your actual secret key
		return []byte("secretKey"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to parse claims"})
	}

	// Set user claims to the context for further use in controllers
	c.Locals("userClaims", token.Claims.(jwt.MapClaims))

	// Continue to the next middleware or handler
	return c.Next()
}

//func VerifyToken(ctx *gin.Context) (interface{}, error) {
//	headerToken := ctx.Request.Header.Get("Authorization")
//	bearer := strings.HasPrefix(headerToken, "Bearer")
//
//	if !bearer {
//		return nil, errors.New("bearer token not found")
//	}
//
//	stringToken := headerToken[7:]
//
//	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
//		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errors.New("failed to get sign token")
//		}
//
//		return []byte(secretKey), nil
//	})
//	if err != nil {
//		return nil, errors.New("invalid token")
//	}
//
//	if _, ok := token.Claims.(jwt.MapClaims); !ok {
//		return nil, errors.New("failed to parse claims")
//	}
//
//	return token.Claims.(jwt.MapClaims), nil
//}
