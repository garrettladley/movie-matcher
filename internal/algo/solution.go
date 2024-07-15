package algo

import (
	"context"
	"movie-matcher/internal/movie"
	"movie-matcher/internal/services/omdb"
	"movie-matcher/internal/services/pref_gen"
	"movie-matcher/internal/set"
	"regexp"
	"slices"
	"strings"
	"sync"
)

var (
	onlyAlphaNumericRegex *regexp.Regexp = regexp.MustCompile("[^a-zA-Z0-9]+")
	orderedRatings                       = set.New(string(movie.RatingG), string(movie.RatingPG), string(movie.RatingPG13), string(movie.RatingR), string(movie.RatingNC17))
)

// Returns the optimal ranking of the given movies for the given people. Authored by Jackson.
func jacksonSolution(ctx context.Context, movies set.OrderedSet[movie.ID], people []pref_gen.Person) set.OrderedSet[movie.ID] {
	scores := make([]struct {
		id    movie.ID
		score int
	}, len(movies.Slice()))

	// Separate goroutine for each movie.
	var wg sync.WaitGroup
	for i, id := range movies.Slice() {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Retrieve movie from OMDb service.
			m, err := omdb.FindMovieById(ctx, string(id))
			if err != nil {
				return
			}

			// Calculate the score for each person
			for _, person := range people {
				if person.Preferences.AfterYear != nil && m.Year >= person.Preferences.AfterYear.Value {
					scores[i].score += int(person.Preferences.AfterYear.Weight)
				}
				if person.Preferences.BeforeYear != nil && m.Year < person.Preferences.BeforeYear.Value {
					scores[i].score += int(person.Preferences.BeforeYear.Weight)
				}
				if person.Preferences.MaximumAgeRating != nil {
					max := slices.Index(orderedRatings.Slice(), string(person.Preferences.MaximumAgeRating.Value))
					actual := slices.Index(orderedRatings.Slice(), m.AgeRating)
					if max != -1 && actual != -1 && actual <= max {
						scores[i].score += int(person.Preferences.MaximumAgeRating.Weight)
					}
				}
				if person.Preferences.ShorterThan != nil && m.Duration < person.Preferences.ShorterThan.Value {
					scores[i].score += int(person.Preferences.ShorterThan.Weight)
				}
				if person.Preferences.FavoriteGenre != nil && slices.Contains(m.Genres, person.Preferences.FavoriteGenre.Value) {
					scores[i].score += int(person.Preferences.FavoriteGenre.Weight)
				}
				if person.Preferences.LeastFavoriteDirector != nil && slices.Contains(m.Directors, person.Preferences.LeastFavoriteDirector.Value) {
					scores[i].score -= int(person.Preferences.LeastFavoriteDirector.Weight)
				}
				if person.Preferences.FavoriteActors != nil {
					for _, actor := range person.Preferences.FavoriteActors.Value {
						if slices.Contains(m.Actors, actor) {
							scores[i].score += int(person.Preferences.FavoriteActors.Weight)
						}
					}
				}
				if person.Preferences.FavoritePlotElements != nil {
					plotElements := strings.Split(onlyAlphaNumericRegex.ReplaceAllString(m.Plot, ""), " ")
					for _, element := range person.Preferences.FavoritePlotElements.Value {
						if slices.Contains(plotElements, element) {
							scores[i].score += int(person.Preferences.FavoritePlotElements.Weight)
						}
					}
				}
				if person.Preferences.MinimumRottenTomatoesScore != nil && m.RottenTomatoesScore >= person.Preferences.MinimumRottenTomatoesScore.Value {
					scores[i].score -= int(person.Preferences.MinimumRottenTomatoesScore.Weight)
				}
			}
		}()
		wg.Done()
	}

	wg.Wait()

	// Sort the results, descending by score
	slices.SortFunc(scores, func(a, b struct {
		id    movie.ID
		score int
	}) int {
		return b.score - a.score
	})

	// Convert scores to a set of IDs
	result := []movie.ID{}
	for _, id := range scores {
		result = append(result, id.id)
	}
	return set.New(result...)
}
