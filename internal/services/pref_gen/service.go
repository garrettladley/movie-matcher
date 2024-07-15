package pref_gen

import (
	"math/rand"
	"movie-matcher/internal/duration"
	"movie-matcher/internal/movie"
	"movie-matcher/internal/utilities"
	"time"
)

type Person struct {
	Name        string      `json:"name"`
	Preferences preferences `json:"preferences"`
}

func GeneratePeople(rand *rand.Rand, n uint) []Person {
	people := make([]Person, 0, n)
	for i := 0; uint(i) < n; i++ {
		people = append(people, Person{
			Name:        names[rand.Intn(len(names))],
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
		prefs.MaximumAgeRating = &preference[movie.Rating]{Value: ratings[rand.Intn(len(ratings))], Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(3) == 0 {
		prefs.ShorterThan = &preference[duration.Duration]{Value: duration.Duration(time.Duration(60+rand.Intn(120)) * time.Minute), Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(2) == 0 {
		prefs.FavoriteGenre = &preference[string]{Value: genres[rand.Intn(len(genres))], Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(3) == 0 {
		prefs.LeastFavoriteDirector = &preference[string]{Value: directors[rand.Intn(len(directors))], Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(2) == 0 {
		prefs.FavoriteActors = &preference[[]string]{Value: utilities.SelectRandom(actors[:], 1+rand.Intn(3)), Weight: uint(5 + rand.Intn(30))}
	}
	if rand.Intn(3) == 0 {
		prefs.FavoritePlotElements = &preference[[]string]{Value: utilities.SelectRandom(plotElements[:], 1+rand.Intn(3)), Weight: uint(5 + rand.Intn(30))}
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
	MaximumAgeRating           *preference[movie.Rating]      `json:"maximumAgeRating(inclusive),omitempty"`
	ShorterThan                *preference[duration.Duration] `json:"shorterThan(exclusive),omitempty"`
	FavoriteGenre              *preference[string]            `json:"favoriteGenre,omitempty"`
	LeastFavoriteDirector      *preference[string]            `json:"leastFavoriteDirector,omitempty"`
	FavoriteActors             *preference[[]string]          `json:"favoriteActors,omitempty"`
	FavoritePlotElements       *preference[[]string]          `json:"favoritePlotElements,omitempty"`
	MinimumRottenTomatoesScore *preference[uint]              `json:"minimumRottenTomatoesScore(inclusive),omitempty"`
}
