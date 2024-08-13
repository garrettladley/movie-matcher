// Package omdb provides functionality for interacting with the OMDb API, a RESTful Open Movie Database
package omdb

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"movie-matcher/internal/duration"
)

type Movie struct {
	IMDbID              string            `json:"imdbID"`
	Title               string            `json:"title"`
	Year                uint              `json:"year"`
	AgeRating           string            `json:"ageRating"`
	Duration            duration.Duration `json:"duration"`
	Genres              []string          `json:"genres"`
	Directors           []string          `json:"directors"`
	Writers             []string          `json:"writers"`
	Actors              []string          `json:"actors"`
	Plot                string            `json:"plot"`
	Languages           []string          `json:"languages"`
	Poster              string            `json:"poster"`
	IMDbScore           uint              `json:"imdbScore"`
	RottenTomatoesScore uint              `json:"rottenTomatoesScore"`
	MetacriticScore     uint              `json:"metacriticScore"`
}

func FindMovieById(ctx context.Context, id string) (Movie, error) {
	res, err := client.query(ctx, params{ID: id, PlotLength: "full", ResultType: "movie", ApiVersion: 1})
	if err != nil {
		return Movie{}, err
	}
	if res.Type != "movie" {
		return Movie{}, fmt.Errorf("failed to find a movie with the ID: %s", id)
	}
	return movieFromResult(res), nil
}

func FindMovieByTitle(ctx context.Context, title string) (Movie, error) {
	res, err := client.query(ctx, params{Title: title, PlotLength: "full", ResultType: "movie", ApiVersion: 1})
	if err != nil {
		return Movie{}, err
	}
	if res.Type != "movie" {
		return Movie{}, fmt.Errorf("failed to find a movie with the title: %s", title)
	}
	return movieFromResult(res), nil
}

func movieFromResult(res result) Movie {
	year, _ := strconv.Atoi(res.Year)
	runtime, _ := time.ParseDuration(fmt.Sprintf("%sm", strings.Split(res.Runtime, " ")[0]))
	var (
		imdbScore           uint
		rottenTomatoesScore uint
		metacriticScore     uint
	)
	for _, rating := range res.Ratings {
		switch rating.Source {
		case "Internet Movie Database":
			strScore, _, found := strings.Cut(rating.Value, "/")
			if !found {
				continue
			}
			score, err := strconv.ParseFloat(strScore, 32)
			if err != nil {
				continue
			}
			imdbScore = uint(score * 10)
		case "Rotten Tomatoes":
			strScore, _, found := strings.Cut(rating.Value, "%")
			if !found {
				continue
			}
			score, err := strconv.ParseUint(strScore, 10, 8)
			if err != nil {
				continue
			}
			rottenTomatoesScore = uint(score)
		case "Metacritic":
			strScore, _, found := strings.Cut(rating.Value, "/")
			if !found {
				continue
			}
			score, err := strconv.ParseUint(strScore, 10, 8)
			if err != nil {
				continue
			}
			metacriticScore = uint(score)
		}
	}
	return Movie{
		IMDbID:              res.IMDbID,
		Title:               res.Title,
		Year:                uint(year),
		AgeRating:           res.Rated,
		Duration:            duration.Duration(runtime),
		Genres:              strings.Split(res.Genre, ", "),
		Directors:           strings.Split(res.Director, ", "),
		Writers:             strings.Split(res.Writer, ", "),
		Actors:              strings.Split(res.Actors, ", "),
		Plot:                res.Plot,
		Languages:           strings.Split(res.Language, ", "),
		Poster:              res.Poster,
		IMDbScore:           imdbScore,
		RottenTomatoesScore: rottenTomatoesScore,
		MetacriticScore:     metacriticScore,
	}
}

var client = &apiClient{
	apiKey: func() string {
		return os.Getenv("OMDB_API_KEY")
	},
}
