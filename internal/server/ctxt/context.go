package ctxt

import (
	"context"

	"movie-matcher/internal/applicant"

	"github.com/gofiber/fiber/v2"
)

type contextKey byte

const (
	contextKeyEmail contextKey = iota
)

func WithEmail(c *fiber.Ctx, email applicant.NUEmail) {
	c.Locals(contextKeyEmail, email)
}

func GetEmail(ctx context.Context) applicant.NUEmail {
	return ctx.Value(contextKeyEmail).(applicant.NUEmail)
}
