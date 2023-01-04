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
	// get firebaseAuth struct from fiber context
	firebaseAuth := utils.GetLocal[*auth.Client](c, "firebaseAuth")
	authorizationToken := c.GetReqHeaders()["Authorisation"]
	// trim header value to just get the id token string
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		// append error string to the e slice
		e = append(e, "id token not available")
		// send JSON response back containing error slice using the base "types.Response" struct
		return c.JSON(types.Response{Error: e})
	}
	//verify token with firebase
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		e = append(e, "invalid token")
		return c.JSON(types.Response{Error: e})
	}
	// set the UUID of the user to the fiber context
	utils.SetLocal[string](c, "uuid", token.UID)
	c.Set("UUID", token.UID)
	return c.Next()
}
