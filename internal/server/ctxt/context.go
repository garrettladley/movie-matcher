package ctxt

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type tokenKey struct{}

func WithToken(c *fiber.Ctx, token uuid.UUID) {
	c.Locals(tokenKey{}, token)
}

func GetToken(ctx context.Context) uuid.UUID {
	return ctx.Value(tokenKey{}).(uuid.UUID)
}
