package ctxt

import (
	"context"

	"movie-matcher/internal/applicant"

	"github.com/gofiber/fiber/v2"
)

type emailKey struct{}

func WithEmail(c *fiber.Ctx, email applicant.NUEmail) {
	c.Locals(emailKey{}, email)
}

func GetEmail(ctx context.Context) applicant.NUEmail {
	return ctx.Value(emailKey{}).(applicant.NUEmail)
}
