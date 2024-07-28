package data

import (
	"time"

	"movie-matcher/internal/views/types"
)

// [[time, value], [time, value], ...]
func Into[T types.Number](timeseries []types.TimePoint[T]) [][]interface{} {
	data := make([][]interface{}, len(timeseries))

	for i, point := range timeseries {
		data[i] = []interface{}{
			point.Time.UnixNano() / int64(time.Millisecond),
			point.Value,
		}
	}

	return data
}
