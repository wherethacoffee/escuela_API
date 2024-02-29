package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wherethacoffe/escuela_API/models"
)

func CreateAccessToken(user models.User) (string, error) {
    secretKey := os.Getenv("SECRET_KEY")

    if secretKey == "" {
	return "", fmt.Errorf("Secret key not set in environment variables")
    }

    myKey := []byte(secretKey)

    claims := jwt.RegisteredClaims{
	Issuer: user.Id.Hex(),
	Subject: user.Username,
	ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    ss, err := token.SignedString(myKey)
    if err != nil {
	return "", err
    }

    return ss, nil
}

func ValidateToken(cookieName string) fiber.Handler {
    secretKey := os.Getenv("SECRET_KEY")

    myKey := []byte(secretKey)

   return func(c *fiber.Ctx) error {
	cookie := c.Cookies(cookieName)
	if cookie == "" {
	    return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
		"error": "Unauthorized",
	    })
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
	return []byte(myKey), nil
   })

   if err != nil || !token.Valid {
	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
	    "error": "Unauthorized",
	})
   }

   return c.Next()
   }
}
