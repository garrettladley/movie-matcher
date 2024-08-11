package handlers

import (
	"movie-matcher/internal/model"
	"movie-matcher/internal/views/types"
)

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
