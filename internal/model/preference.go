package model

type Preference[T any] struct {
	Name   string
	Value  T
	Weight int
}
