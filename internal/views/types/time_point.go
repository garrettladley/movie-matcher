package types

import (
	"time"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

type TimePoint[T Number] struct {
	Value T
	Time  time.Time
}
