package handlers

import (
	"time"

	"movie-matcher/internal/views/status"
	"movie-matcher/internal/views/types"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Status(c *fiber.Ctx) error {
	values := []int{6500, 6418, 6456, 6526, 6356, 6456}
	categories := []string{"01 February", "02 February", "03 February", "04 February", "05 February", "06 February"}

	dataPoints := make([]types.TimePoint[int], len(values))
	for i, value := range values {
		dataPoints[i] = types.TimePoint[int]{
			Value: value,
			Time:  parseDate(categories[i]),
		}
	}

	return into(c, status.Index(dataPoints))
}

func parseDate(dateStr string) time.Time {
	layout := "02 January"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		panic(err)
	}
	// Set year to a fixed value (e.g., current year) if needed
	return t
}
