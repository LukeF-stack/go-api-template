package middleware

import (
	"context"
	"example/bookAPI/internal/server/types"
	"example/bookAPI/internal/server/utils"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware(c *fiber.Ctx) error {
	var e []string
	firebaseAuth := utils.GetLocal[*auth.Client](c, "firebaseAuth")
	authorizationToken := c.GetReqHeaders()["Authorisation"]
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		e = append(e, "id token not available")
		return c.JSON(types.Response{Error: e})
	}
	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		e = append(e, "invalid token")
		return c.JSON(types.Response{Error: e})
	}
	c.Set("UUID", token.UID)
	return c.Next()
}
