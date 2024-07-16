package algo

import (
	"context"
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"sync/atomic"

	"movie-matcher/internal/movie"
	"movie-matcher/internal/ordered_set"
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/services/pref_gen"
	"movie-matcher/internal/utilities"
)

type Prompt struct {
	Movies ordered_set.OrderedSet[movie.ID] `json:"movies"`
	People []pref_gen.Person                `json:"people"`
}

type Service struct {
	client *omdb.CachedClient
}

func NewService(client *omdb.CachedClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) Generate(rand *rand.Rand) Prompt {
	return Prompt{
		Movies: ordered_set.New(utilities.SelectRandom(movie.Catalog, 10)...),
		People: pref_gen.GeneratePeople(rand, 30),
	}
}

func (s *Service) Check(ctx context.Context, expected ordered_set.OrderedSet[movie.ID], actual ordered_set.OrderedSet[movie.ID]) int {
	return ordered_set.Distance(expected, actual)
}

var orderedRatings = ordered_set.New(pref_gen.Ratings...)

type movieScore struct {
	id    movie.ID
	score uint32
}

func (s *Service) Solution(ctx context.Context, movies ordered_set.OrderedSet[movie.ID], people []pref_gen.Person) (ordered_set.OrderedSet[movie.ID], error) {
	scores := make([]movieScore, len(movies.Slice()))
	var wg sync.WaitGroup
	errChan := make(chan error, len(movies.Slice()))

	for i, id := range movies.Slice() {
		wg.Add(1)
		go func(i int, id movie.ID) {
			defer wg.Done()
			score, err := s.calculateScoreForMovie(ctx, id, people)
			if err != nil {
				errChan <- fmt.Errorf("failed to calculate score for movie %s: %w", id, err)
				return
			}
			scores[i] = score
		}(i, id)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return ordered_set.OrderedSet[movie.ID]{}, <-errChan
	}

	sortScores(scores)
	return extractMovieIDs(scores), nil
}

func (s *Service) calculateScoreForMovie(ctx context.Context, id movie.ID, people []pref_gen.Person) (movieScore, error) {
	om, err := s.client.FindMovieById(ctx, string(id))
	if err != nil {
		return movieScore{id: id}, fmt.Errorf("error finding movie by ID: %w", err)
	}

	m := movie.FromOMDB(om)

	var totalScore uint32
	var wg sync.WaitGroup

	for _, person := range people {
		wg.Add(1)
		go func(person pref_gen.Person) {
			defer wg.Done()
			personScore := calculatePersonScore(m, person)
			atomic.AddUint32(&totalScore, personScore)
		}(person)
	}

	wg.Wait()

	return movieScore{id: id, score: totalScore}, nil
}

func calculatePersonScore(m movie.Movie, person pref_gen.Person) uint32 {
	var score uint

	if person.Preferences.AfterYear != nil && m.Year >= person.Preferences.AfterYear.Value {
		score += person.Preferences.AfterYear.Weight
	}
	if person.Preferences.BeforeYear != nil && m.Year < person.Preferences.BeforeYear.Value {
		score += person.Preferences.BeforeYear.Weight
	}

	if person.Preferences.MaximumAgeRating != nil {
		max := slices.Index(orderedRatings.Slice(), string(person.Preferences.MaximumAgeRating.Value))
		actual := slices.Index(orderedRatings.Slice(), m.AgeRating)
		if max != -1 && actual != -1 && actual <= max {
			score += person.Preferences.MaximumAgeRating.Weight
		}
	}

	if person.Preferences.ShorterThan != nil && m.Duration < person.Preferences.ShorterThan.Value {
		score += person.Preferences.ShorterThan.Weight
	}

	if person.Preferences.FavoriteGenre != nil && slices.Contains(m.Genres, person.Preferences.FavoriteGenre.Value) {
		score += person.Preferences.FavoriteGenre.Weight
	}

	if person.Preferences.LeastFavoriteDirector != nil && slices.Contains(m.Directors, person.Preferences.LeastFavoriteDirector.Value) {
		score -= person.Preferences.LeastFavoriteDirector.Weight
	}

	if person.Preferences.FavoriteActors != nil {
		score += person.Preferences.FavoriteActors.Weight * utilities.IntersectionCardinality(m.Actors, person.Preferences.FavoriteActors.Value)
	}

	if person.Preferences.FavoritePlotElements != nil {
		score += person.Preferences.FavoritePlotElements.Weight * utilities.IntersectionCardinality(m.Plot, person.Preferences.FavoritePlotElements.Value)
	}

	if person.Preferences.MinimumRottenTomatoesScore != nil && m.RottenTomatoesScore >= person.Preferences.MinimumRottenTomatoesScore.Value {
		score -= person.Preferences.MinimumRottenTomatoesScore.Weight
	}

	return uint32(score)
}

func sortScores(scores []movieScore) {
	slices.SortFunc(scores, func(a, b movieScore) int {
		return int(b.score - a.score)
	})
}

func extractMovieIDs(scores []movieScore) ordered_set.OrderedSet[movie.ID] {
	result := make([]movie.ID, 0, len(scores))
	for _, s := range scores {
		result = append(result, s.id)
	}
	return ordered_set.New(result...)
}
