package ctxt

import (
	"context"

	"movie-matcher/internal/applicant"

	"github.com/gofiber/fiber/v2"
)

type contextKey byte

const (
	contextKeyNUID contextKey = iota
)

func WithNUID(c *fiber.Ctx, nuid applicant.NUID) {
	c.Locals(contextKeyNUID, nuid)
}

func GetNUID(ctx context.Context) applicant.NUID {
	return ctx.Value(contextKeyNUID).(applicant.NUID)
}
