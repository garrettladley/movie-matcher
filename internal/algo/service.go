package algo

import (
	"context"
	"math/rand"
	"movie-matcher/internal/movie"
	"movie-matcher/internal/services/pref_gen"
	"movie-matcher/internal/set"
	"movie-matcher/internal/utilities"
)

type Prompt struct {
	Movies set.OrderedSet[movie.ID] `json:"movies"`
	People []pref_gen.Person        `json:"people"`
}

func Generate(rand *rand.Rand) Prompt {
	movies := utilities.SelectRandom(movie.Catalog, 10)
	movieIDs := []movie.ID{}
	for _, movie := range movies {
		movieIDs = append(movieIDs, movie.ID)
	}
	return Prompt{
		// TODO: Movies list
		Movies: set.New(movieIDs...),
		People: pref_gen.GeneratePeople(rand, 30),
	}
}

func Check(prompt Prompt, actual set.OrderedSet[movie.ID]) uint {
	expected := jacksonSolution(context.Background(), prompt.Movies, prompt.People)
	return set.Distance(expected, actual)
}
