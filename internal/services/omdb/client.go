package omdb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type apiClient struct {
	apiKey func() string
}

func (c *apiClient) query(ctx context.Context, params params) (result, error) {
	requestUrl, err := buildQueryUrl(c.apiKey(), params)
	if err != nil {
		return result{}, fmt.Errorf("failed to build request URL: %w", err)
	}

	agent := fiber.Get(requestUrl.String())
	agent.JSONDecoder(go_json.Unmarshal)

	resultCh := make(chan struct {
		res result
		err error
	})

	go func() {
		var res result
		statusCode, _, errs := agent.Struct(&res)
		if len(errs) > 0 {
			resultCh <- struct {
				res result
				err error
			}{result{}, fmt.Errorf("failed to make request: %w", errors.Join(errs...))}
			return
		}

		if statusCode != http.StatusOK {
			resultCh <- struct {
				res result
				err error
			}{result{}, fmt.Errorf("received status code: %d", statusCode)}
			return
		}

		if res.Response == "False" {
			resultCh <- struct {
				res result
				err error
			}{result{}, fmt.Errorf("received an error from API: %s", res.Error)}
			return
		}

		resultCh <- struct {
			res result
			err error
		}{res, nil}
	}()

	select {
	case <-ctx.Done():
		return result{}, ctx.Err()
	case res := <-resultCh:
		return res.res, res.err
	}
}

type params struct {
	ID         string
	Title      string
	Year       string
	PlotLength string
	ResultType string
	ApiVersion int
}

type result struct {
	Response string
	Error    string
	Title    string
	Year     string
	Rated    string
	Released string
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Language string
	Country  string
	Awards   string
	Poster   string
	Ratings  []struct {
		Source string
		Value  string
	}
	Metascore  string
	IMDbRating string `json:"imdbRating"`
	IMDbVotes  string `json:"imdbVotes"`
	IMDbID     string `json:"imdbID"`
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
}

func buildQueryUrl(apiKey string, params params) (*url.URL, error) {
	queryParams := url.Values{}
	queryParams.Set("apikey", apiKey)

	if (params.ID == "" && params.Title == "") || (params.ID != "" && params.Title != "") {
		return &url.URL{}, fmt.Errorf("must specify either an id or a title, but not both")
	}
	if params.ID != "" {
		queryParams.Set("i", params.ID)
	}
	if params.Title != "" {
		queryParams.Set("t", params.Title)
	}

	if params.ResultType == "movie" || params.ResultType == "series" || params.ResultType == "episode" {
		queryParams.Set("type", params.ResultType)
	} else if params.ResultType != "" {
		return &url.URL{}, fmt.Errorf("invalid result type specified: %s", params.ResultType)
	}

	if params.Year != "" {
		queryParams.Set("y", params.Year)
	}

	if params.PlotLength == "short" || params.PlotLength == "full" {
		queryParams.Set("plot", params.PlotLength)
	} else if params.PlotLength != "" {
		return &url.URL{}, fmt.Errorf("invalid plot length specified: %s", params.PlotLength)
	}

	if params.ApiVersion != 0 {
		queryParams.Set("v", strconv.Itoa(params.ApiVersion))
	}

	requestUrl, _ := url.Parse("http://www.omdbapi.com/")

	requestUrl.RawQuery = queryParams.Encode()
	return requestUrl, nil
}
