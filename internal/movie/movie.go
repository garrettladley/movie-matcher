package movie

import (
	"fmt"
	"math/rand"
	"movie-matcher/internal/duration"
	"movie-matcher/internal/services/omdb"
	"time"
)

type Movie struct {
	Title               string
	ID                  ID
	Year                uint
	AgeRating           string
	Duration            duration.Duration
	Languages           []string
	Genres              []string
	Directors           []string
	Writers             []string
	Actors              []string
	Plot                []string
	IMDbScore           uint
	RottenTomatoesScore uint
	MetacriticScore     uint
}

type MovieDisplayDetails struct {
	PosterURL 		string
	TimeRemaining 	*int 	`json:"timeRemaining,omitempty"`
}

type MovieDisplay struct {
	Movie
	MovieDisplayDetails
}

type ID string // IMDb id

func FromOMDB(omdbMovie omdb.Movie) Movie {
	return Movie{
		Title:               omdbMovie.Title,
		ID:                  ID(omdbMovie.IMDbID),
		Year:                omdbMovie.Year,
		AgeRating:           omdbMovie.AgeRating,
		Duration:            omdbMovie.Duration,
		Languages:           omdbMovie.Languages,
		Genres:              omdbMovie.Genres,
		Directors:           omdbMovie.Directors,
		Writers:             omdbMovie.Writers,
		Actors:              omdbMovie.Actors,
		Plot:                omdbMovie.Plot,
		IMDbScore:           omdbMovie.IMDbScore,
		RottenTomatoesScore: omdbMovie.RottenTomatoesScore,
		MetacriticScore:     omdbMovie.MetacriticScore,
	}
}

// Convert a movie to a display movie:
func MovieToDisplay(movie Movie, curWatching bool) (*MovieDisplay, error) {
	displayDetails, keyFound := CatalogDisplay[movie.ID]
	if !keyFound {
		return nil, fmt.Errorf("movie with ID %s not found in frontend catalog", movie.ID)
	}

	// compute a random number of minutes remaining for the movie
	if curWatching {
		randMinutes := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(120) + 1
		displayDetails.TimeRemaining = &randMinutes
	}

	return &MovieDisplay{
		Movie: movie,
		MovieDisplayDetails: displayDetails,
	}, nil
}
