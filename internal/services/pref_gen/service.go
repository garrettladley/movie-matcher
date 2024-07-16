package pref_gen

import (
	"math/rand"
	"time"

	"movie-matcher/internal/duration"
	"movie-matcher/internal/utilities"
)

type Person struct {
	Name        string      `json:"name"`
	Preferences preferences `json:"preferences"`
}

func GeneratePeople(rand *rand.Rand, n uint) []Person {
	people := make([]Person, 0, n)
	in := int(n)
	for index := 0; index < in; index++ {
		people = append(people, Person{
			Name:        Names[rand.Intn(len(Names))],
			Preferences: generatePreferences(rand),
		})
	}
	return people
}

func generatePreferences(rand *rand.Rand) preferences {
	prefs := preferences{}
	if rand.Intn(3) == 0 {
		prefs.AfterYear = &preference[uint]{Value: uint(1900 + rand.Intn(150)), Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(3) == 0 {
		prefs.BeforeYear = &preference[uint]{Value: uint(1900 + rand.Intn(150)), Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(3) == 0 {
		prefs.MaximumAgeRating = &preference[string]{Value: Ratings[rand.Intn(len(Ratings))], Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(3) == 0 {
		prefs.ShorterThan = &preference[duration.Duration]{Value: duration.Duration(time.Duration(60+rand.Intn(120)) * time.Minute), Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(2) == 0 {
		prefs.FavoriteGenre = &preference[string]{Value: Genres[rand.Intn(len(Genres))], Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(3) == 0 {
		prefs.LeastFavoriteDirector = &preference[string]{Value: Directors[rand.Intn(len(Directors))], Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(2) == 0 {
		prefs.FavoriteActors = &preference[[]string]{Value: utilities.SelectRandom(Actors[:], 1+rand.Intn(3)), Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(3) == 0 {
		prefs.FavoritePlotElements = &preference[[]string]{Value: utilities.SelectRandom(PlotElements[:], 1+rand.Intn(3)), Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(2) == 0 {
		prefs.MinimumRottenTomatoesScore = &preference[uint]{Value: uint(20 + rand.Intn(81)), Weight: uint(5 + rand.Intn(30))}
	}
	return prefs
}

type preference[T any] struct {
	Value  T    `json:"value"`
	Weight uint `json:"weight"`
}

type preferences struct {
	AfterYear                  *preference[uint]              `json:"afterYear(inclusive),omitempty"`
	BeforeYear                 *preference[uint]              `json:"beforeYear(exclusive),omitempty"`
	MaximumAgeRating           *preference[string]            `json:"maximumAgeRating(inclusive),omitempty"`
	ShorterThan                *preference[duration.Duration] `json:"shorterThan(exclusive),omitempty"`
	FavoriteGenre              *preference[string]            `json:"favoriteGenre,omitempty"`
	LeastFavoriteDirector      *preference[string]            `json:"leastFavoriteDirector,omitempty"`
	FavoriteActors             *preference[[]string]          `json:"favoriteActors,omitempty"`
	FavoritePlotElements       *preference[[]string]          `json:"favoritePlotElements,omitempty"`
	MinimumRottenTomatoesScore *preference[uint]              `json:"minimumRottenTomatoesScore(inclusive),omitempty"`
}
