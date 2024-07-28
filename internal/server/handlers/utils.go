package handlers

import (
	"movie-matcher/internal/model"
	"movie-matcher/internal/views/types"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func into(c *fiber.Ctx, component templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}

func intoTimePoints(submissions []model.Submission) []types.TimePoint[int] {
	dataPoints := make([]types.TimePoint[int], len(submissions))

	for i, submission := range submissions {
		dataPoints[i] = types.TimePoint[int]{
			Value: submission.Score,
			Time:  submission.Time,
		}
	}

	return dataPoints
}
