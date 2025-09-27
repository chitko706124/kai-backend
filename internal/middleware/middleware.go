package middleware

import (
	"net/http"

	"github.com/LouisFernando1204/kai-backend.git/dto"
	"github.com/LouisFernando1204/kai-backend.git/internal/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected(conf *config.Config) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(conf.Jwt.Key)},
		TokenLookup:  "header:Authorization",
		AuthScheme:   "Bearer",
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			if userID, ok := claims["user_id"].(string); ok {
				c.Locals("user_id", userID)
			} else {
				return c.Status(http.StatusUnauthorized).JSON(dto.CreateResponseWithoutData(http.StatusUnauthorized, "Invalid token claims"))
			}

			return c.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {

	errMsg := err.Error()

	switch {
	case errMsg == "Missing or malformed JWT":
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, "Missing or malformed JWT. Please provide Authorization header with Bearer token"))
	case errMsg == "token is malformed":
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, "Token is malformed"))
	case errMsg == "token is expired":
		return c.Status(http.StatusUnauthorized).JSON(dto.CreateResponseWithoutData(http.StatusUnauthorized, "Token is expired"))
	case errMsg == "signature is invalid":
		return c.Status(http.StatusUnauthorized).JSON(dto.CreateResponseWithoutData(http.StatusUnauthorized, "Token signature is invalid - JWT key mismatch"))
	default:
		return c.Status(http.StatusUnauthorized).JSON(dto.CreateResponseWithoutData(http.StatusUnauthorized, "JWT Error: "+errMsg))
	}
}
