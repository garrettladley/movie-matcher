package omdb

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/storage/memory/v2"
)

type CachedClient struct {
	*memory.Storage
}

func NewCachedClient() *CachedClient {
	return &CachedClient{memory.New()}
}

func (c *CachedClient) FindMovieById(ctx context.Context, id string) (Movie, error) {
	key := fmt.Sprintf("%d%s", idPrefix, id)
	movie, ok := c.checkCache(key)
	if ok {
		return movie, nil
	}

	movie, err := FindMovieById(ctx, id)
	if err != nil {
		return Movie{}, err
	}

	marshalledMovie, err := go_json.Marshal(movie)
	if err != nil {
		return Movie{}, err
	}

	if err := c.Set(key, marshalledMovie, 30*time.Minute); err != nil {
		slog.Error("error caching movie", "id", id, "error", err)
	}

	return movie, nil
}

func (c *CachedClient) FindMovieByTitle(ctx context.Context, title string) (Movie, error) {
	key := fmt.Sprintf("%d%s", titlePrefix, title)
	movie, ok := c.checkCache(key)
	if ok {
		return movie, nil
	}

	movie, err := FindMovieByTitle(ctx, title)
	if err != nil {
		return Movie{}, err
	}

	marshalledMovie, err := go_json.Marshal(movie)
	if err != nil {
		return Movie{}, err
	}

	if err := c.Set(key, marshalledMovie, 30*time.Minute); err != nil {
		slog.Error("error caching movie", "title", title, "error", err)
	}

	return movie, nil
}

type cachePrefix byte

const (
	idPrefix    cachePrefix = 0
	titlePrefix cachePrefix = 1
)

func (c *CachedClient) checkCache(key string) (Movie, bool) {
	movieBytes, err := c.Get(key)
	if err != nil {
		return Movie{}, false
	}

	var movie Movie
	if err := go_json.Unmarshal(movieBytes, &movie); err != nil {
		slog.Error("error unmarshalling movie", "key", key, "error", err)
		return Movie{}, false
	}

	return movie, true
}
