package omdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type apiClient struct {
	apiKey     func() string
	httpClient *http.Client
}

func (c *apiClient) query(ctx context.Context, params params) (result, error) {
	requestUrl, err := buildQueryUrl(c.apiKey(), params)
	if err != nil {
		return result{}, fmt.Errorf("failed to build request URL: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return result{}, fmt.Errorf("failed to build request: %w", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return result{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return result{}, fmt.Errorf("received status code: %d", response.StatusCode)
	}

	var res result
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return result{}, fmt.Errorf("failed to decode result: %w", err)
	}
	if res.Response == "False" {
		return result{}, fmt.Errorf("received an error from API: %s", res.Error)
	}
	return res, nil
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

	if params.PlotLength == "short" || params.PlotLength == "long" {
		queryParams.Set("plot", params.PlotLength)
	} else if params.PlotLength != "" {
		return &url.URL{}, fmt.Errorf("invalid plot length specified: %s", params.PlotLength)
	}

	if params.ApiVersion != 0 {
		queryParams.Set("v", strconv.Itoa(params.ApiVersion))
	}

	requestUrl, err := url.Parse("http://www.omdbapi.com/")
	if err != nil {
		return &url.URL{}, fmt.Errorf("failed to parse base URL: %w", err)
	}

	requestUrl.RawQuery = queryParams.Encode()
	return requestUrl, nil
}
