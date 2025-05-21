package identity

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiberMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        if c.Path() == "/identity-health-check" {
            checkCode := c.Query("check_code")
            return c.SendString(c.Method() + " " + c.Path() + " " + checkCode)
        }
        c.Locals("userID", "anonymous")
        return c.Next()
    }
}